# Tonic

<img align="right" src="assets/tonic.jpg" width="313" alt="Tonic Logo">
<img src="https://godoc.org/github.com/golang/gddo?status.svg" alt="GoDoc">
<img src="https://github.com/duolok/tonic/workflows/build/badge.svg" alt="Build Status">

Tonic is a delightful, minimalistic, and easy-to-use framework for building customizable 2D art applications. It's designed with simplicity and flexibility in mind, offering a clean way to create generative art using intuitive, Go-based APIs.

<p>
    <img src="https://example.com/tonic-example.gif" width="100%" alt="Tonic Art Example">
</p>

Whether you're a seasoned developer or an artist, Tonic is here to help you build, tweak, and generate unique pieces of art with minimal effort. Its lightweight design makes it perfect for hobby projects or integrating creative visuals into larger applications.

To get started, check out the tutorial below, the [examples][examples], the [docs][docs], or some of the common [resources](#resources-we-use-with-tonic).

## Features

- Intuitive Go-based API for creating generative art
- High-performance rendering with minimal dependencies
- Wide variety of artistic engines and tools to choose from
- Customizable color palettes, backgrounds, and line properties
- Supports exporting artwork as PNG files

## Tutorial

Tonic is designed to be approachable and simple. Below is a quick tutorial that assumes basic knowledge of Go.

By the way, the non-annotated source code for this program is available [on GitHub][tut-source].

### The Model

The model represents your canvas and the engine driving your artwork.

```go
type model struct {
    engines  []generativeart.Engine // Generative art engines
    selected map[int]struct{}       // Selected art pieces
}
```

