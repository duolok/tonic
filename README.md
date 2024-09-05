# Tonic

<img align="right" src="assets/tonic.jpg" width="313" alt="Tonic Logo">
<a href="https://pkg.go.dev/github.com/duolok/tonic"><img src="https://godoc.org/github.com/golang/gddo?status.svg" alt="GoDoc"></a>
<a href="https://github.com/duolok/tonic/actions"><img src="https://github.com/duolok/tonic/workflows/go/badge.svg" alt="Build Status"></a>

Tonic is a delightful, minimalistic, and easy-to-use framework for building customizable 2D art applications. It's designed with simplicity and flexibility in mind, offering a clean way to create generative art using intuitive, Go-based APIs.

<p>
    <img src="https://example.com/tonic-example.gif" width="100%" alt="Tonic Art Example">
</p>

Whether you're a seasoned developer or an artist, Tonic is here to help you build, tweak, and generate unique pieces of art with minimal effort. Its lightweight design makes it perfect for hobby projects or integrating creative visuals into larger applications.

## Features

- Intuitive Go-based API for creating generative art
- High-performance rendering with minimal dependencies
- Wide variety of artistic engines and tools to choose from
- Customizable color palettes, backgrounds, and line properties
- Supports exporting artwork as PNG files

## Getting Started

### Prerequisites

Tonic requires [Go](https://go.dev/) version 1.18 or higher.

### Installing Tonic

To install Tonic, use `go get` to add the package to your project:

```sh
go get -u github.com/duolok/tonic
```

Example Usage
A simple example to create an artwork using the "circles" generator:

```go
package main

import (
  "github.com/gin-gonic/gin"
  "github.com/duolok/tonic/generator"
  "net/http"
)

func main() {
  r := gin.Default()
  r.GET("/art/circles", func(c *gin.Context) {
    file := generator.DrawOne("circles")
    c.Header("Content-Type", "image/png")
    c.File(file)
  })
  r.Run()
}
```

To run the code, use:

```sh
go run main.go
```
### Then, visit 0.0.0.0:8080/art/circles to see the generated artwork!

### Documentation
See the full API documentation.

### Contributing
We welcome contributions from the community! Please see CONTRIBUTING.md for details on submitting patches and the contribution workflow.

### Resources 
- Bubbles: Common UI components for terminal applications.
- Lip Gloss: Tools for styling and formatting terminal UIs.
- Termenv: Advanced ANSI styling.
