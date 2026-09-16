# clay
[![GoDoc](https://godoc.org/github.com/TotallyGamerJet/clay?status.svg)](https://godoc.org/github.com/TotallyGamerJet/clay)

This is a Go port of the C layout library ([Clay](https://github.com/nicbarker/clay)).

## Goals

* Be entirely Go (no Cgo)
* Idiomatic Go public API
* No [unsafe](https://pkg.go.dev/unsafe) in the public API
* Be fast

## Renderers
Clay includes multiple prebuilt renderers:

* Ebitengine - [hajimehoshi/ebiten](https://github.com/hajimehoshi/ebiten)
* Raylib - [gen2brain/raylib-go](https://github.com/gen2brain/raylib-go)
* SDL2 - [veandco/go-sdl2](https://github.com/veandco/go-sdl2)
* SDL3 - [Zyko0/go-sdl3](https://github.com/Zyko0/go-sdl3)
* Software - [golang.org/x/image](https://golang.org/x/image)

## Updating clay

Clay's C source (`clay.h`) is compiled to WebAssembly with clang and translated to Go with
[wasm2go](https://github.com/ncruces/wasm2go). The Go types and bindings in `clay.go` are
generated from clang's AST of `clay.h`. To update clay, replace `clay.h` and run:

```shell
go generate
```

This requires a clang that can target WebAssembly, like [wasi-sdk](https://github.com/WebAssembly/wasi-sdk).
It is looked up in `tools/wasi-sdk`, `$WASI_SDK_PATH`, or `PATH`; set `CLANG` to override.

## License

This project is governed by the MIT license. See LICENSE for full description.
