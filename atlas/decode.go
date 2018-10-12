package atlas

import (
	"fmt"
	"image"
	"io"

	"github.com/pkg/errors"
	"github.com/xackery/tmx/model"
	"github.com/xackery/tmx/pb"
	"golang.org/x/image/draw"
)

// Decoder decodes image files to compacted image file
type Decoder struct {
	r   io.Reader
	t   *pb.TileSet
	img image.Image
}

// NewDecoder creates a new decoder
func NewDecoder(r io.Reader, t *pb.TileSet) (d *Decoder) {
	d = &Decoder{
		r: r,
		t: t,
	}
	return
}

// Decode reads a TMX xml and outputs to a protobuf file
func (d *Decoder) Decode(a *Atlas) (err error) {
	a.tileMap = make(map[int64]int64)
	a.tiles = make(map[int64]*image.RGBA)
	d.img, _, err = image.Decode(d.r)
	if err != nil {
		err = errors.Wrap(err, "image decode")
		return
	}
	tileSize := image.Rect(0, 0, int(d.t.TileWidth), int(d.t.TileHeight))
	a.tileMap = make(map[int64]int64)
	oldIndex := int64(0)
	var offset image.Point

	total := d.t.TileCount
	oldPercent := int64(-1)

	for offset.Y = d.img.Bounds().Min.Y; offset.Y < d.img.Bounds().Max.Y; offset.Y += int(d.t.TileHeight) {
		for offset.X = d.img.Bounds().Min.X; offset.X < d.img.Bounds().Max.X; offset.X += int(d.t.TileWidth) {
			tile := image.NewRGBA(tileSize)
			if model.IsVerbose() {
				percent := int64(float64(float64(oldIndex)/float64(total)) * float64(100))
				if oldPercent != percent {
					fmt.Printf("%d%%s, ", percent)
					oldPercent = percent
				}
			}
			draw.Draw(tile, tileSize, d.img, offset, draw.Src)
			newIndex := a.AppendUniqueThread(tile)
			a.tileMap[oldIndex] = newIndex
			oldIndex++
		}
	}

	//need to remove hard code
	sheetWidth := 1024
	img := image.NewRGBA(image.Rect(0, 0, sheetWidth, sheetWidth))

	//may not be needed
	tileSize.Max.X += 60
	tileSize.Max.Y += 60

	offset = image.ZP
	//fmt.Println(len(a.tiles), "total tiles")
	for _, tile := range a.tiles {
		offset.X += int(d.t.TileWidth)
		if offset.X > sheetWidth {
			offset.Y += int(d.t.TileHeight)
			offset.X = 0
		}
		if offset.Y > sheetWidth {
			fmt.Println("exceeded size")
			return
		}
		if tile == nil {
			err = fmt.Errorf("tile nil malfunction")
			return
		}
		//fmt.Printf("%d,%d|", offset.X, offset.Y)

		draw.Copy(img, offset, tile, tileSize, draw.Src, nil)
	}
	a.img = img
	return
}
