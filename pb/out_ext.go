package pb

import (
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

	return data
}
