[![GoDoc](https://godoc.org/github.com/xackery/goeq?status.svg)](https://godoc.org/github.com/xackery/goeq)
[![Go Report Card](https://goreportcard.com/badge/github.com/xackery/tmx)](https://goreportcard.com/report/github.com/xackery/tmx) [![Build Status](https://travis-ci.org/xackery/tmx.svg)](https://travis-ci.org/Xackery/tmx.svg?branch=master) [![Coverage Status](https://coveralls.io/repos/github/xackery/tmx/badge.svg?branch=master)](https://coveralls.io/github/xackery/tmx?branch=master)

# tmx - Tiled loader and packer

This project is not yet released. You can talk about it [here](https://discord.gg/dF55WRZ)

TMX takes a [Tiled](https://www.mapeditor.org/) .tmx file and transforms it into an optimized asset for games. It parses the tmx base file, any external tile sets (tsx), along with all image assets and repacks them to optimize for in game usage. The output files are [protobuf](https://developers.google.com/protocol-buffers/) defined.

## Use as an executable

* [download from releases](https://github.com/xackery/tmx/releases) or `go get -u github.com/xackery/tmx`
* `usage: tmx source_file destination_file`
* supported output: .bin, .data, .json, .xml, .yml

### Use tool as a library

You can import tmx into your go projects and create custom flows.

```go
package main

import (
  "context"
  "fmt"

  "github.com/xackery/tmx/client"
)

func main() {
  ctx := context.Background()
  c, err := client.New(ctx)
  if err != nil {
    panic(err)
  }
  tmx, err := c.LoadFile(ctx, "file.tmx")
  if err != nil {
    panic(err)
  }
  fmt.Println(tmx)
}
```

## Planned Features

* Documentation on how to load protobuf serialized .bin and .data files

* Finish GID indexing and mirror/flip detection