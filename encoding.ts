import pl from "npm:nodejs-polars";

export function oneHotEncoding(dataframe) {
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
