package tmx

import (
	"fmt"

	"github.com/xackery/tmx/model"
	"github.com/xackery/tmx/pb"
)

//walkRoot is the root index nodes
func walkTMXRoot(m *pb.Map, nodes []Node) (err error) {
	for _, n := range nodes {
		switch n.XMLName.Local {
		case "map":
			for _, attr := range n.Attrs {
				switch attr.Name.Local {
				case "version":
					m.Version = attr.Value
				case "tiledversion":
					m.TiledVersion = attr.Value
				case "orientation":
					m.Orientation = attr.Value
				case "renderorder":
					m.RenderOrder = attr.Value
				case "width":
					m.Width = model.ToInt64(attr.Value)
				case "height":
					m.Height = model.ToInt64(attr.Value)
				case "tilewidth":
					m.TileWidth = model.ToInt64(attr.Value)
				case "tileheight":
					m.TileHeight = model.ToInt64(attr.Value)
				case "infinite":
					m.Infinite = model.ToInt64(attr.Value)
				case "hexsidelength":
					m.HexSideLength = model.ToInt64(attr.Value)
				case "staggeraxis":
					m.StaggerAxis = attr.Value
				case "backgroundcolor":
					m.BackgroundColor = attr.Value
				case "nextlayerid":
					m.NextLayerId = attr.Value
				case "nextobjectid":
					m.NextObjectId = attr.Value
				default:
					fmt.Println("unknown attribute inside", n.XMLName.Local, "name:", attr.Name.Local, "value:", attr.Value)
				}
			}
			err = walkTMXMap(m, n.Nodes)
			if err != nil {
				return
			}
		default:
			fmt.Println("unknown element inside root:", n.XMLName.Local)
		}
	}
	return
}

func walkTMXMap(m *pb.Map, nodes []Node) (err error) {
	for _, n := range nodes {
		//Can contain: <properties>, <tileset>, <layer>, <objectgroup>, <imagelayer>, <group> (since 1.0)
		switch n.XMLName.Local {
		case "layer":
			l := &pb.Layer{}
			for _, attr := range n.Attrs {
				switch attr.Name.Local {
				case "id":
					l.Id = model.ToInt64(attr.Value)
				case "name":
					l.Name = attr.Value
				case "x":
					l.X = model.ToInt64(attr.Value)
				case "y":
					l.Y = model.ToInt64(attr.Value)
				case "height":
					l.Height = model.ToInt64(attr.Value)
				case "width":
					l.Width = model.ToInt64(attr.Value)
				case "opacity":
					l.Opacity = model.ToFloat32(attr.Value)
				case "Visible":
					l.Visible = model.ToBool(attr.Value)
				default:
					fmt.Println("unknown attribute inside", n.XMLName.Local, "name:", attr.Name.Local, "value:", attr.Value)
				}
			}
			err = walkTMXLayer(l, n.Nodes)
			if err != nil {
				return
			}
			m.Layers = append(m.Layers, l)
		case "tileset":
			t := &pb.TileSet{}
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
				default:
					fmt.Println("unknown attribute inside", n.XMLName.Local, "name:", attr.Name.Local, "value:", attr.Value)
				}

			}
			m.TileSets = append(m.TileSets, t)
		default:
			fmt.Println("unknown element inside map:", n.XMLName.Local)
		}
	}
	return
}

func walkTMXLayer(l *pb.Layer, nodes []Node) (err error) {
	for _, n := range nodes {
		//Can contain: <properties>, <data>
		switch n.XMLName.Local {
		case "data":
			d := &pb.Data{}
			for _, attr := range n.Attrs {
				switch attr.Name.Local {
				case "encoding":
					d.Encoding = attr.Value
				case "compression":
					d.Compression = attr.Value
				default:
					fmt.Println("unknown attribute inside", n.XMLName.Local, "name:", attr.Name.Local, "value:", attr.Value)
				}
			}
			d.RawData = n.Content
			if err != nil {
				return
			}
			err = walkTMXData(d, n.Nodes)
			if err != nil {
				return
			}
			l.Data = d
		default:
			fmt.Println("unknown element inside layer:", n.XMLName.Local)
		}
	}
	return
}

func walkTMXData(d *pb.Data, nodes []Node) (err error) {
	for _, n := range nodes {
		//Can contain: <tile>, <chunk>
		switch n.XMLName.Local {
		case "tile":
			t := &pb.DataTile{}
			for _, attr := range n.Attrs {
				switch attr.Name.Local {
				case "gid":
					t.Gid = model.ToInt64(attr.Value)
				default:
					fmt.Println("unknown attribute inside", n.XMLName.Local, "name:", attr.Name.Local, "value:", attr.Value)
				}
			}
			d.DataTiles = append(d.DataTiles, t)
		case "chunk":
			c := &pb.Chunk{}
			for _, attr := range n.Attrs {
				switch attr.Name.Local {
				case "x":
					c.X = model.ToInt64(attr.Value)
				case "y":
					c.Y = model.ToInt64(attr.Value)
				case "width":
					c.Width = model.ToInt64(attr.Value)
				case "height":
					c.Height = model.ToInt64(attr.Value)
				default:
					fmt.Println("unknown attribute inside", n.XMLName.Local, "name:", attr.Name.Local, "value:", attr.Value)
				}
			}
			d.Chunks = append(d.Chunks, c)
		default:
			fmt.Println("unknown element inside layer:", n.XMLName.Local)
		}
	}
	return
}
