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
* SDL2 - [veandco/go-sdl2](https://github.com/veandco/go-sdl2)
* SDL3 - [Zyko0/go-sdl3](https://github.com/Zyko0/go-sdl3)
* Software - [golang.org/x/image](https://golang.org/x/image)

## Generate clay.go

Everything in `clay.go` is generated from the main project's `clay.h` file using [CxGo](https://github.com/gotranspile/cxgo).
Just run the command below to regenerate it.
```shell
go tool cxgo
```

## License

This project is governed by the MIT license. See LICENSE for full description.
