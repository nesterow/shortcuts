import pl from "npm:nodejs-polars";

export function residuals(x: pl.Expr, y: pl.Expr): pl.Expr {
  const xM = x.minus(x.mean());
  const yM = y.minus(y.mean());
  const xMSQ = xM.pow(2);
  const beta = xM.dot(yM).div(xMSQ.sum());
  return yM.minus(beta.mul(xM));
}

export const fillzero = (
  value = 0.0001,
) => (pl.all().replaceStrict(0, value, pl.all()));

export const ScaleExpr: pl.Expr = (pl.all().minus(pl.all().min())).div(
  pl.all().max().minus(pl.all().min()),
);
export const StdNormExpr: pl.Expr = pl.all().minus(pl.all().mean()).div(
  pl.all().std(),
);
