import pl from "npm:nodejs-polars";

export function oneHotEncoding(dataframe: pl.DataFrame): pl.DataFrame {
  let df = pl.DataFrame();
  for (const columnName of dataframe.columns) {
    const column = dataframe[columnName];
    if (!column.isNumeric()) {
      df = df.hstack(column.toDummies());
    } else {
      df = df.hstack(dataframe.select(columnName));
    }
  }
  return df;
}

export function polynomialTransform(
  dataframe: pl.DataFrame,
  degree = 2,
  interaction_only = false,
  include_bias = true,
): pl.DataFrame {
  let polyRecords: number[][] = [];
  dataframe.map((X: number[]) => {
    polyRecords.push(
      polynomialFeatures(X, degree, interaction_only, include_bias),
    );
  });
  return pl.readRecords(polyRecords);
}

export function polynomialFeatures(
  X: number[],
  degree = 2,
  interaction_only = false,
  include_bias = true,
): number[] {
  let features = [...X];
  let prev_chunk = [...X];
  const indices = Array.from({ length: X.length }, (_, i) => i);
  for (let d = 1; d < degree; d++) {
    const new_chunk: any[] = [];
    for (let i = 0; i < (interaction_only ? X.length - d : X.length); i++) {
      const v = X[i];
      const next_index = new_chunk.length;
      for (let j = i + (interaction_only ? 1 : 0); j < prev_chunk.length; j++) {
        new_chunk.push(v * prev_chunk[j]);
      }
      indices[i] = next_index;
    }
    features = features.concat(new_chunk);
    prev_chunk = new_chunk;
  }
  if (include_bias) {
    features.unshift(1);
  }
  return features;
}

/**
 * Add rows at given interval, use average to fill values.
 * Usage:
 * ```ts
 * let df = augmentMeanForward("price", df, 100);
 * ```
 * @param feature
 * @param df
 * @param interval
 */
export function augmentMeanForward(
  feature: string,
  df: pl.DataFrame,
  interval = 100,
) {
  let sorted = df.sort(feature);
  let featIdx = sorted.findIdxByName(feature);
  let result = sorted.head(1);
  for (let i = 0; i < sorted.height; i++) {
    let p1 = sorted.row(i).at(featIdx);
    let k = (i + 1) % sorted.height;
    let p2 = sorted.row(k).at(featIdx);
    if (p2 - p1 > interval) {
      for (let j = 0; j < Math.round((p2 - p1) / interval) - 1; j++) {
        result = pl.concat([
          result,
          pl.concat([
            result.tail(1),
            sorted.slice({ offset: k, length: 1 }),
            sorted.head(1).shift(-1),
          ])
            .fillNull("mean")
            .tail(1),
        ]);
      }
    } else {
      result = pl.concat([result, sorted.slice(1, i)]);
    }
  }
  return result;
}
