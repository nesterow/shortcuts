import pl from "npm:nodejs-polars";

type DfSplit = {
  trainX: pl.DataFrame;
  trainY: pl.DataFrame;
  testX: pl.DataFrame;
  testY: pl.DataFrame;
  size: number;
};

export function sliceK(
  df: pl.DataFrame,
  size: number,
  k: number,
  ...yFeatures: string[]
): DfSplit[] {
  let testSize = Math.round(df.shape.height * size);
  while (testSize % k !== 0) {
    testSize -= 1;
  }
  if (df.shape.height / testSize < k) {
    throw new Error(
      `k value is too large, max k value is ${df.shape.height / testSize}`,
    );
  }
  let trainSize = df.shape.height - testSize;
  let result: DfSplit[] = [];
  let data = df;
  for (let i = 0; i < k; i++) {
    let [train, test] = [data.head(trainSize), data.tail(testSize)];
    let [trainY, testY] = [
      train.select(...yFeatures),
      test.select(...yFeatures),
    ];
    let [trainX, testX] = [train.drop(...yFeatures), test.drop(...yFeatures)];
    result.push({
      trainX,
      trainY,
      testX,
      testY,
      size,
    });
    data = pl.concat([test, train]);
  }
  return result;
}

export function trainTestSplit(
  df: pl.DataFrame,
  size: number,
  shuffle = true,
  ...yFeatures: string[]
) {
  let data = shuffle ? df.sample(df.height - 1) : df;
  const result = sliceK(data, size, 1, ...yFeatures);
  return result[0];
}

export function kFold(
  df: pl.DataFrame,
  k: number,
  shuffle = true,
  ...yFeatures: string[]
): DfSplit[] {
  let data = shuffle ? df.sample(df.height - 1) : df;
  return sliceK(data, 1 / k, k, ...yFeatures);
}
