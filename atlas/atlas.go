package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/pkg/errors"
	"golang.org/x/image/draw"
)

// Atlas represents an image atlas
type Atlas struct {
	tiles      map[int]*image.RGBA
	tileWidth  int
	tileHeight int
}

// NewAtlas returns a new instance of atlas
func NewAtlas(tileWidth, tileHeight int) (a *Atlas) {
	a = &Atlas{
		tiles:      make(map[int]*image.RGBA),
		tileWidth:  tileWidth,
		tileHeight: tileHeight,
	}
	return
}

// AppendUnique will append if unique, otherwise returns existing index
func (a *Atlas) AppendUnique(img *image.RGBA) (index int) {
	var tile *image.RGBA
	for index, tile = range a.tiles {
		isMatch := true
		for y := img.Bounds().Min.Y; isMatch && y < img.Bounds().Max.Y; y++ {
			for x := img.Bounds().Min.X; isMatch && x < img.Bounds().Max.X; x++ {
				//fmt.Println(tile.At(x, y), img.At(x, y))
				if tile.At(x, y) != img.At(x, y) {
					isMatch = false
				}
			}
		}
		if isMatch {
			return
		}
	}

	index = len(a.tiles)
	a.tiles[index] = img
	return
}

// Image returns the map as an image
func (a *Atlas) Image() (img *image.RGBA) {
	sheetWidth := 1024
	img = image.NewRGBA(image.Rect(0, 0, sheetWidth, sheetWidth))
	tileSize := image.Rect(0, 0, a.tileWidth+60, a.tileHeight+60)
	offset := image.ZP
	fmt.Println(len(a.tiles), "total tiles")
	for _, tile := range a.tiles {
		offset.X += a.tileWidth
		if offset.X > sheetWidth {
			offset.Y += a.tileHeight
			offset.X = 0
		}
		if offset.Y > sheetWidth {
			fmt.Println("exceeded size")
			return
		}
		if tile == nil {

		}
		//fmt.Printf("%d,%d|", offset.X, offset.Y)

		draw.Copy(img, offset, tile, tileSize, draw.Src, nil)
	}
	fmt.Println("final offset", offset)
	return
}

func main() {
	err := run()
	if err != nil {
		fmt.Println("error:", err.Error())
	}
}

func run() (err error) {
	var src, dst string
	args := os.Args
	if len(args) < 3 {
		usage()
		return
	}

	if len(args) > 1 {
		src = args[len(args)-2]
	}
	if len(args) > 2 {
		dst = args[len(args)-1]
	}

	if len(src) == 0 || len(dst) == 0 {
		usage()
		return
	}

	tileWidth := 8
	tileHeight := 8
	r, err := os.Open(src)
	if err != nil {
		err = errors.Wrap(err, "failed to load")
		return
	}
	img, ext, err := image.Decode(r)
	if err != nil {
		err = errors.Wrap(err, "failed to decode")
		return
	}
	tileSize := image.Rect(0, 0, tileWidth, tileHeight)

	a := NewAtlas(tileWidth, tileHeight)

	var offset image.Point
	for offset.Y = img.Bounds().Min.Y; offset.Y < img.Bounds().Max.Y; offset.Y += tileHeight {
		for offset.X = img.Bounds().Min.X; offset.X < img.Bounds().Max.X; offset.X += tileWidth {
			tile := image.NewRGBA(tileSize)
			draw.Draw(tile, tileSize, img, offset, draw.Src)
			a.AppendUnique(tile)
			//fmt.Println(index)
		}
	}

	var f *os.File
	switch ext {
	case "png":
		f, err = os.Create(dst)
		if err != nil {
			err = errors.Wrap(err, "failed to open out file")
			return
		}
		e := png.Encoder{
			CompressionLevel: png.BestCompression,
		}
		err = e.Encode(f, a.Image())
		if err != nil {
			err = errors.Wrap(err, "failed to encode")
			return
		}
		f.Close()
	default:
		err = fmt.Errorf("unknown file type: %s", ext)
		return
	}
	return
}

func usage() (err error) {
	fmt.Println("usage: atlas source_file target_file")
	return
}
