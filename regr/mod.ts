import "../lib/wasm_tinygo.js";

// @ts-expect-error: no types
const go = new Go();

const code =
  await (await fetch(import.meta.url.replace("/mod.ts", "/mod.wasm")))
    .arrayBuffer();

const wasmMmodule = await WebAssembly.instantiate(code, go.importObject);
const wasm = wasmMmodule.instance;

go.run(wasm);

// @ts-ignore: no types
const _exports =  __InitRegrExports(_exports) as Record<string, (...args: unknown[]) => unknown>;

export default _exports;
