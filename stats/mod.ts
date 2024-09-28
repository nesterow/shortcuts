import "../lib/wasm.js";

const go = new Go();
const code =
  await (await fetch(import.meta.url.replace("/mod.ts", "/mod.wasm")))
    .arrayBuffer();
const wasmMmodule = await WebAssembly.instantiate(code, go.importObject);
const wasm = wasmMmodule.instance;
go.run(wasm);
export default wasm.exports;
