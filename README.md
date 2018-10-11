# tmx

Tiled TMX and TSX file support

## Usage

`tmx source_file destination_file`

### Use as a library

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