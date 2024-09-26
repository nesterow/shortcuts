import pl from "npm:nodejs-polars";

export function trainTestSplit(df: pl.DataFrame, size, ...yFeatures: string[]) {
  let shuffle = df.sample(df.height - 1);
  let testSize = Math.round(shuffle.shape.height * size);
  let trainSize = shuffle.shape.height - testSize;
  let [train, test] = [shuffle.head(trainSize), shuffle.tail(testSize)];
  let [trainY, testY] = [train.select(...yFeatures), test.select(...yFeatures)];
  let [trainX, testX] = [train.drop(...yFeatures), test.drop(...yFeatures)];
  return {
    trainX,
    trainY,
    testX,
    testY,
  };
}
