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
	m         *pb.Map
	r         io.Reader
	t         *pb.TileSet
	img       image.Image
	usedTiles map[uint32]bool
}

// NewDecoder creates a new decoder
func NewDecoder(r io.Reader, m *pb.Map, t *pb.TileSet, usedTiles map[uint32]bool) (d *Decoder) {
	d = &Decoder{
		r:         r,
		m:         m,
		t:         t,
		usedTiles: usedTiles,
	}
	return
}

// Decode reads a TMX xml and outputs to a protobuf file
func (d *Decoder) Decode(a *Atlas) (err error) {
	if a.tileMap == nil {
		a.tileMap = make(map[uint32]*model.GID)
		a.tiles = make(map[int]*image.RGBA)
		a.tileWidth = d.t.TileWidth
		a.tileHeight = d.t.TileHeight
	}

	if a.tileWidth != d.t.TileWidth || a.tileHeight != d.t.TileHeight {
		err = fmt.Errorf("all tilesets must use the same with/height for this program to work")
		return
	}

	tw := a.tileWidth
	th := a.tileHeight
	d.img, _, err = image.Decode(d.r)
	if err != nil {
		err = errors.Wrap(err, "image decode")
		return
	}
	tileSize := image.Rect(0, 0, int(tw), int(th))
	oldIndex := uint32(d.t.FirstGid)
	//fmt.Println("firstgid", oldIndex)
	var offset image.Point

	var gid *model.GID
	var ok bool
	newTileCountBefore := len(a.tiles)
	for offset.Y = d.img.Bounds().Min.Y; offset.Y < d.img.Bounds().Max.Y; offset.Y += int(th) {
		for offset.X = d.img.Bounds().Min.X; offset.X < d.img.Bounds().Max.X; offset.X += int(tw) {
			a.oldTileCount++
			if _, ok = d.usedTiles[oldIndex]; !ok { //skip unused tiles
				oldIndex++
				continue
			}
			tile := image.NewRGBA(tileSize)

			draw.Draw(tile, tileSize, d.img, offset, draw.Src)
			gid = a.AppendUnique(tile)

			//fmt.Println("mapping", oldIndex, gid.Index())
			if oldIndex < 5 {
				fmt.Println(gid)
			}
			a.tileMap[oldIndex] = gid
			oldIndex++
		}
	}
	a.newTileCount += len(a.tiles) - newTileCountBefore

	return
}
