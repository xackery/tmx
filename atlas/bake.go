package atlas

import (
	"fmt"
	"image"

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

			//fmt.Println(i, oldGid.Index(), oldGid.HRead(), oldGid.VRead(), oldGid.DRead(), oldGid.ValueRead(), newGid.ValueRead(), newGid.Index())
			//err = newGid.RotationUpdate(oldGid.RotationRead())
			//if err != nil {
			//	return
			//}
			//fmt.Println("remapped", oldGid.ValueRead(), oldGid.Index(), oldGid.RotationRead(), "to", newGid.ValueRead(), newGid.Index(), "rotation", newGid.RotationRead())
			if newGid == nil {
				continue
			}
			//newGid.RotationUpdate(oldGid.RotationRead())
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
	return
}
