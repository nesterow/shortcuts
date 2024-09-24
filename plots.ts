import * as Plot from "npm:@observablehq/plot";
import { DOMParser, SVGElement } from "npm:linkedom";

const defaultPlotSettings = {
    grid: true,
    margin: 50,
    style: {
        backgroundColor: "#fff"
    }
}

/**
 * Configure default plot settings
 * @param options 
 */
export function configurePlots(options: any) {
    Object.assign(defaultPlotSettings, options)
}


/**
 * Draw side-by-side plots 
 * Example:
 * ```ts
 * const plt = sideBySidePlot({
 *   x: ['enginesize', 'horsepower', 'citympg', 'highwaympg'],
 *   y: ['price'],
 *   marks: [
 *     (x, y) => Plot.dot(data, {x, y, fill: "species"}),
 *     (x, y) => Plot.linearRegressionY(data, {x, y, stroke: "red"}),
 *   ],
 *   cols: 2
 * })
 * @param x List of x-axis targets
 * @param y List of y-axis targets
 * @param marks List of plot callbacks
 * @param cols Number of columns
 */
export function sideBySidePlot(opts: { 
    x: string[], 
    y: string[],
    marks: any[]
    cols: number,
    options,
}) {
    const document = new DOMParser().parseFromString(`<!DOCTYPE html><html></html>`, "text/html");
    const imgTags: string[] = []
    for (const xTarget of opts.x) {
        for (const yTarget of opts.y) {
            const plt = Plot.plot({
                ...(opts.options ?? defaultPlotSettings),
                marks: opts.marks.map(fn => fn(xTarget, yTarget)),
                document
            })
            plt.setAttribute('xmlns', 'http://www.w3.org/2000/svg')
            const svgUrl = `data:image/svg+xml;base64,${btoa(unescape(encodeURIComponent(plt)))}`
            imgTags.push(`<img title="${xTarget} / ${yTarget}" src='${svgUrl}'>`)
        }
    }
    const output = `
          <section style="display:grid;grid-template-columns: repeat(${opts.cols ?? 2}, 1fr);">
              ${imgTags.join('')}
          </section>
    `
    return {
        [Symbol.for("Jupyter.display")]: () => ({
            "text/html": output
        })
    }
}