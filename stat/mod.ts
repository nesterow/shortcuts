import "../lib/wasm.js";
import type { Stat } from "./types.ts";

// @ts-expect-error: no types
const go = new Go();

const code =
  await (await fetch(import.meta.url.replace("/mod.ts", "/mod.wasm")))
    .arrayBuffer();

const wasmMmodule = await WebAssembly.instantiate(code, go.importObject);
const wasm = wasmMmodule.instance;

go.run(wasm);

const _exports = {} as Record<string, (...args: unknown[]) => unknown> & Stat;

// @ts-ignore: no types
__InitStatExports(_exports);

export default _exports;
