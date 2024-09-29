import "../lib/wasm_exec.js";

// @ts-expect-error: no types
const go = new Go();

const code =
  await (await fetch(import.meta.url.replace("/mod.ts", "/mod.wasm")))
    .arrayBuffer();

const wasmMmodule = await WebAssembly.instantiate(code, go.importObject);
const wasm = wasmMmodule.instance;

go.run(wasm);

const _exports = {} as Record<string, (...args: unknown[]) => unknown>;

// @ts-ignore: no types
__InitPlotExports(_exports);

for (const key in _exports) {
  const draw = _exports[key];
  const drawKey = "Draw" + key;
  _exports[drawKey] = (...args: unknown[]) => {
    const data = "data:image/png;base64," + draw(...args);
    return {
      [Symbol.for("Jupyter.display")]: () => ({
        "text/markdown": `![name](${data})`,
      }),
    };
  };
}

export default _exports;
