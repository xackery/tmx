package atlas

import (
	"fmt"
	"image"
	"image/color"
	"strings"

	"github.com/xackery/tmx/model"
	"github.com/xackery/tmx/pb"
	"golang.org/x/image/draw"
)

// Bake returns the internal image
func (a *Atlas) Bake(m *pb.Map) (img *image.RGBA, err error) {

	//now we remap the map data with the new tilmap
	for i := range m.Layers {
		for j := 0; j < len(m.Layers[i].Data.DataTiles); j++ {
			d := m.Layers[i].Data.DataTiles[j]
			if d.Gid == 0 {
				continue
			}
			oldGid := model.NewGID(d.Gid)

			newGid := a.tileMap[model.Index(d.Gid)]
			if oldGid.HRead() && newGid.HRead() {
				newGid.HUpdate(false)
			}
			newGid.HUpdate(oldGid.HRead())
			if oldGid.VRead() && newGid.VRead() {
				newGid.VUpdate(false)
			}
			newGid.VUpdate(oldGid.VRead())
			if oldGid.DRead() && newGid.DRead() {
				newGid.DUpdate(false)
			}
			newGid.DUpdate(oldGid.DRead())

			if newGid == nil {
				continue
			}

			m.Layers[i].Data.DataTiles[j].Gid = newGid.ValueRead()

		}
	}

	//need to remove hard code
	sheetWidth := 2048
	sheetHeight := sheetWidth
	//sheetWidth := 512
	img = image.NewRGBA(image.Rect(0, 0, sheetWidth, sheetHeight))
	//may not be needed
	//a.tileWidth += 60
	//a.tileHeight += 60
	tileSize := image.Rect(0, 0, int(a.tileWidth), int(a.tileHeight))
	offset := image.ZP
	//fmt.Println(len(a.tiles), "total tiles")
	for i := 0; i < len(a.tiles); i++ {
		tile := a.tiles[i]

		if tile == nil {
			err = fmt.Errorf("tile nil malfunction")
			return
		}
		//fmt.Printf("%d,%d|", offset.X, offset.Y)

		draw.Copy(img, offset, tile, tileSize, draw.Src, nil)

		offset.X += int(a.tileWidth)
		if offset.X > sheetWidth {
			offset.Y += int(a.tileHeight)
			offset.X = 0
			fmt.Println("new row", i)
		}
		if offset.Y > sheetHeight {
			fmt.Println("exceeded size")
			return
		}
	}
	a.img = img

	tileset := "tileset"
	if len(m.Layers) > 0 {
		tileset = "tilesets"
	}
	fmt.Println(tileset, "reduced from", a.oldTileCount, "to", a.newTileCount, "total tiles")

	//build collider
	if len(m.Layers) < 1 {
		return
	}
	empty := color.RGBA{0, 0, 0, 0}
	for i := range m.Layers[0].Data.DataTiles {
		isCollider := false
		for _, l := range m.Layers {
			if strings.Contains(strings.ToLower(l.Name), "bg") {
				continue
			}
			t := l.Data.DataTiles[i]
			id := model.Index(t.Gid)
			if id == 0 {
				continue
			}
			img := a.tiles[int(id)]
			if img == nil {
				continue
			}
			for y := 0; y < img.Bounds().Max.Y && !isCollider; y++ {
				for x := 0; x < img.Bounds().Max.X && !isCollider; x++ {
					//fmt.Println(img.At(x, y))
					if img.At(x, y) != empty {
						isCollider = true
					}
				}
			}
			if isCollider {
				break
			}
		}
		m.Colliders = append(m.Colliders, &pb.Collider{IsCollider: isCollider})
	}
	return
}
