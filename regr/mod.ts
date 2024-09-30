import "../lib/wasm_tinygo.js";

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
__InitRegrExports(_exports);

export default _exports;
