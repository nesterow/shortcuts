import pl from "npm:nodejs-polars"

export function residuals(x: pl.Expr, y: pl.Expr): pl.Expr {
    const xM = x.minus(x.mean())
    const yM = y.minus(y.mean())
    const xMSQ = xM.pow(2)
    const beta = xM.dot(yM).div(xMSQ.sum())
    return yM.minus(beta.mul(xM))
}