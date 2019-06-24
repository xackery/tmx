package pb

import (
	fmt "fmt"
	"image"

	"github.com/xackery/egui"
)

// ToEGUI converts an outmap to egui mapdata
func (m *OutMap) ToEGUI() egui.MapData {
	data := egui.MapData{
		Source:     m.Source,
		Width:      m.Width,
		Height:     m.Height,
		TileWidth:  m.TileWidth,
		TileHeight: m.TileHeight,
		TileCount:  m.TileCount,
	}
	for i := range m.Layers {
		l := m.Layers[i]
		nl := egui.MapLayer{
			Name:    l.Name,
			Opacity: l.Opacity,
		}
		for j := range l.Tiles {
			t := l.Tiles[j]
			nt := egui.MapTile{
				GID: t.Gid,
			}
			nl.Tiles = append(nl.Tiles, nt)
		}
		data.Layers = append(data.Layers, nl)
	}
	for i := range m.Colliders {
		c := m.Colliders[i]
		nc := egui.MapCollider{
			IsCollider: c.IsCollider,
			Cost:       c.Cost,
		}
		data.Colliders = append(data.Colliders, nc)
	}

	data.TileSheetWidth = data.Width * data.TileWidth
	data.TileSheetHeight = data.Height * data.TileHeight

	tileFrames := []image.Rectangle{}
	tileFrames = append(tileFrames, image.Rect(0, 0, 0, 0))
	tileCount := 0
	var x, y int64

	for y < data.TileSheetHeight { //tileCount < int(om.TileCount) {
		tileFrames = append(tileFrames, image.Rect(int(x), int(y), int(x+data.TileWidth), int(y+data.TileHeight)))
		fmt.Println(int(x), int(y), int(x+data.TileWidth), int(y+data.TileHeight))
		x += data.TileWidth
		if x > data.TileSheetWidth {
			x = 0
			y += data.TileHeight
		}
		//fmt.Println(tileCount, x, y)
		tileCount++
	}
	data.TileFrames = tileFrames
	return data
}
