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
 * Adds missing rows at given interval, uses mean of previous and next value.
 * Example for one feature: [1, 2, 4, 5] -> [1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5]
 * @param feature
 * @param df
 * @param bin
 */
export function augmentMeanForward(
  feature: string,
  df: pl.DataFrame,
  interval = 100,
) {
  let sorted = df.sort(feature);
  let result = sorted.head(0);
  let n: null | number = null;
  for (let i = 0; i < sorted.height - 1; i++) {
    let p1 = n ?? sorted.row(i).at(-1);
    let p2 = sorted.row(i + 1).at(-1);
    if (p2 - p1 > interval) {
      let avg = (p1 + p2) / 2;
      result = pl.concat([
        result,
        pl.concat([result.tail(2), sorted.slice({ offset: i + 1, length: 2 })])
          .shift(-1)
          .fillNull("mean")
          .tail(1),
      ]);
      if (p2 - avg > interval) {
        i--;
        n = avg;
        continue;
      }
      result = pl.concat([result, sorted.slice(1, i)]);
      n = null;
    }
  }
  return result;
}
