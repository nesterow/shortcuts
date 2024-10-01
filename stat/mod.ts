import "../lib/wasm_tinygo.js";
import { Stat } from "./types.ts";

// @ts-expect-error: no types
const go = new Go();

const code =
  await (await fetch(import.meta.url.replace("/mod.ts", "/mod.wasm")))
    .arrayBuffer();

const wasmMmodule = await WebAssembly.instantiate(code, go.importObject);
const wasm = wasmMmodule.instance;

go.run(wasm);

// @ts-ignore: no types
const _exports = __InitStatExports() as Record<string, (...args: unknown[]) => unknown> & Stat;


export default _exports;
