package tsx

import (
	"fmt"

	"github.com/xackery/tmx/model"
	"github.com/xackery/tmx/pb"
)

func walkTSXRoot(t *pb.TileSet, nodes []Node) (err error) {
	for _, n := range nodes {
		switch n.XMLName.Local {
		case "tileset":
			for _, attr := range n.Attrs {
				switch attr.Name.Local {
				case "firstgid":
					t.FirstGid = model.ToInt64(attr.Value)
				case "source":
					t.Source = attr.Value
				case "name":
					t.Name = attr.Value
				case "tilewidth":
					t.TileWidth = model.ToInt64(attr.Value)
				case "tileheight":
					t.TileHeight = model.ToInt64(attr.Value)
				case "spacing":
					t.Spacing = model.ToInt64(attr.Value)
				case "margin":
					t.Margin = model.ToInt64(attr.Value)
				case "tilecount":
					t.TileCount = model.ToInt64(attr.Value)
				case "columns":
					t.ColumnCount = model.ToInt64(attr.Value)
				case "version":
					t.Version = attr.Value
				case "tiledversion":
					t.TiledVersion = attr.Value
				default:
					fmt.Println("tmx?attr", n.XMLName.Local, "name:", attr.Name.Local, "value:", attr.Value)
				}

			}
			err = walkTSXTileSet(t, n.Nodes)
			if err != nil {
				return
			}
		default:
			fmt.Println("tsx>?element@root?", n.XMLName.Local)
		}
	}
	return
}

func walkTSXTileSet(t *pb.TileSet, nodes []Node) (err error) {
	prefix := "tsx"
	scope := "tileset"
	for _, n := range nodes {
		//Can contain: <properties>, <tileset>, <layer>, <objectgroup>, <imagelayer>, <group> (since 1.0)
		switch n.XMLName.Local {
		case "image":
			i := &pb.Image{}
			for _, attr := range n.Attrs {
				switch attr.Name.Local {
				case "format":
					i.Format = attr.Value
				case "source":
					i.Source = attr.Value
				case "trans":
					i.Trans = attr.Value
				case "width":
					i.Width = model.ToInt64(attr.Value)
				case "height":
					i.Height = model.ToInt64(attr.Value)
				default:
					fmt.Printf("%s>%s>%s@attr?%s:%s\n", prefix, scope, n.XMLName.Local, attr.Name.Local, attr.Value)
				}
			}
			t.Image = i
		case "tile":
			ti := &pb.Tile{}
			for _, attr := range n.Attrs {
				switch attr.Name.Local {
				case "id":
					ti.Id = model.ToInt64(attr.Value)
				case "type":
					ti.Type = attr.Value
				case "terrain":
					ti.Terrain = attr.Value
				case "probability":
					ti.Probability = model.ToFloat32(attr.Value)
				default:
					fmt.Printf("%s>%s>%s@attr?%s:%s\n", prefix, scope, n.XMLName.Local, attr.Name.Local, attr.Value)
				}
			}
			t.Tiles = append(t.Tiles, ti)
		default:
			fmt.Printf("%s>%s>%s@element?\n", prefix, scope, n.XMLName.Local)
		}
	}
	return
}
