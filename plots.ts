import * as _Plot from "npm:@observablehq/plot";
import { DOMParser, SVGElement } from "npm:linkedom";
import * as _vega from "npm:vega";
import * as _lite from "npm:vega-lite";

const defaultPlotSettings = {
  grid: true,
  margin: 50,
  style: {
    backgroundColor: "#fff",
  },
};

/**
 * Configure default plot settings
 * @param options
 */
export function configurePlots(options: any) {
  Object.assign(defaultPlotSettings, options);
}

export const document = new DOMParser().parseFromString(
  `<!DOCTYPE html><html></html>`,
  "text/html",
);

export const Plot = _Plot;
export const vega = _vega;
export const vegalite = _lite;

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
  x: string[];
  y: string[];
  marks: any[];
  cols: number;
  options;
}) {
  const imgTags: string[] = [];
  for (const xTarget of opts.x) {
    for (const yTarget of opts.y) {
      const plt = Plot.plot({
        ...(opts.options ?? defaultPlotSettings),
        marks: opts.marks.map((fn) => fn(xTarget, yTarget)),
        document,
      });
      plt.setAttribute("xmlns", "http://www.w3.org/2000/svg");
      const svgUrl = `data:image/svg+xml;base64,${
        btoa(unescape(encodeURIComponent(plt)))
      }`;
      imgTags.push(`<img title="${xTarget} / ${yTarget}" src='${svgUrl}'>`);
    }
  }
  const output = `
          <section style="display:grid;grid-template-columns: repeat(${
    opts.cols ?? 2
  }, 1fr);">
              ${imgTags.join("")}
          </section>
    `;
  return {
    [Symbol.for("Jupyter.display")]: () => ({
      "text/html": output,
    }),
  };
}

/**
 * Histogram plot
 *
 * @param data
 * @param x
 * @param opts
 * @returns Plot
 */
export function histPlot(
  data: any[],
  x = "column",
  opts = { options: null, fn: "proportion" },
) {
  return Plot.plot({
    ...(opts.options ?? defaultPlotSettings),
    y: { grid: true },
    marks: [
      Plot.ruleY([0]),
      Plot.ruleX([0]),
      Plot.rectY(data, Plot.binX({ y: opts.fn ?? "count" }, { x: x })),
    ],
    document,
  });
}

export function oneBoxPlot(
  data: any[],
  y = "column",
  opts = { options: null, box: null },
) {
  return Plot.plot({
    ...(opts.options ?? defaultPlotSettings),
    y: { grid: true },
    marks: [
      Plot.ruleY([0]),
      Plot.boxY(data, { y, ...(opts.box ?? {}) }),
    ],
    document,
  });
}

export async function quantilePlotSVG(
  data: any[],
  x = "column",
  opts = { width: 500 },
) {
  const spec = {
    data: { values: data },
    width: opts.width,
    "transform": [
      {
        "quantile": "price",
        "as": ["prob", "value"],
      },
      {
        "calculate": "quantileNormal(datum.prob)",
        "as": "norm",
      },
    ],
    "layer": [
      {
        "mark": { type: "circle", size: 80 },
        "encoding": {
          "x": {
            "field": "norm",
            "type": "quantitative",
            "title": "Theoretical Quantiles→",
          },
          "y": {
            "field": "value",
            "type": "quantitative",
            "title": "Ordered Values→",
          },
        },
      },
      {
        mark: { type: "line", color: "red" },
        transform: [
          {
            "regression": "value",
            "on": "norm",
          },
        ],
        "encoding": {
          "x": { "field": "norm", "type": "quantitative" },
          "y": { "field": "value", "type": "quantitative" },
        },
      },
    ],
  };
  let vegaspec = vegalite.compile(spec).spec;
  var view = new vega.View(vega.parse(vegaspec), { renderer: "none" });

  return await view.toSVG();
}

export function quantilePlot(data: any[], x = "column", opts = { width: 500 }) {
  return {
    [Symbol.for("Jupyter.display")]: async () => ({
      "text/html": await quantilePlotSVG(data, x, opts),
    }),
  };
}

const svgDataUrl = (plt) =>
  `data:image/svg+xml;base64,${btoa(unescape(encodeURIComponent(plt)))}`;

export function threeChart(data: any[], x = "column", opts = { width: 800 }) {
  const hist = histPlot(data, x);
  hist.setAttribute("xmlns", "http://www.w3.org/2000/svg");
  const box = oneBoxPlot(data, x);
  box.setAttribute("xmlns", "http://www.w3.org/2000/svg");
  const qq = quantilePlotSVG(data, x);
  return {
    [Symbol.for("Jupyter.display")]: async () => ({
      "text/html": `
            <section style="display:flex;flex-direction: column; gap: 1em;">
              <div style="gap:0.5em; display:flex; flex-direction: row; border: 1px solid black;">
                  <img src="${svgDataUrl(hist)}" style="width: 100%;">
                  <img src="${svgDataUrl(box)}" style="width: 65%">
              </div>
              <div style="padding:1em; border: 1px solid black;">
                  <img src="${svgDataUrl(await qq)}" style="width: 100%">
              </div>
            </section>
            `,
    }),
  };
}
