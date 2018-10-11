package client

import (
	"fmt"

	"github.com/xackery/tmx/pb"
)

//walkRoot is the root index nodes
func walkTMXRoot(m *pb.Map, nodes []Node) {
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
					m.Width = toInt64(attr.Value)
				case "height":
					m.Height = toInt64(attr.Value)
				case "tilewidth":
					m.TileWidth = toInt64(attr.Value)
				case "tileheight":
					m.TileHeight = toInt64(attr.Value)
				case "infinite":
					m.Infinite = toInt64(attr.Value)
				case "hexsidelength":
					m.HexSideLength = toInt64(attr.Value)
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
			walkTMXMap(m, n.Nodes)
		default:
			fmt.Println("unknown element inside root:", n.XMLName.Local)
		}
	}
}

func walkTMXMap(m *pb.Map, nodes []Node) {
	for _, n := range nodes {
		//Can contain: <properties>, <tileset>, <layer>, <objectgroup>, <imagelayer>, <group> (since 1.0)
		switch n.XMLName.Local {
		case "layer":
			l := &pb.Layer{}
			for _, attr := range n.Attrs {
				switch attr.Name.Local {
				case "id":
					l.Id = toInt64(attr.Value)
				case "name":
					l.Name = attr.Value
				case "x":
					l.X = toInt64(attr.Value)
				case "y":
					l.Y = toInt64(attr.Value)
				case "height":
					l.Height = toInt64(attr.Value)
				case "width":
					l.Width = toInt64(attr.Value)
				case "opacity":
					l.Opacity = toFloat32(attr.Value)
				case "Visible":
					l.Visible = toBool(attr.Value)
				default:
					fmt.Println("unknown attribute inside", n.XMLName.Local, "name:", attr.Name.Local, "value:", attr.Value)
				}
			}
			walkTMXLayer(l, n.Nodes)
			m.Layers = append(m.Layers, l)
		case "tileset":
			t := &pb.TileSet{}
			for _, attr := range n.Attrs {
				switch attr.Name.Local {
				case "firstgid":
					t.FirstGid = toInt64(attr.Value)
				case "source":
					t.Source = attr.Value
				case "name":
					t.Name = attr.Value
				case "tilewidth":
					t.TileWidth = toInt64(attr.Value)
				case "tileheight":
					t.TileHeight = toInt64(attr.Value)
				case "spacing":
					t.Spacing = toInt64(attr.Value)
				case "margin":
					t.Margin = toInt64(attr.Value)
				case "tilecount":
					t.TileCount = toInt64(attr.Value)
				case "columns":
					t.ColumnCount = toInt64(attr.Value)
				default:
					fmt.Println("unknown attribute inside", n.XMLName.Local, "name:", attr.Name.Local, "value:", attr.Value)
				}

			}
			m.TileSets = append(m.TileSets, t)
		default:
			fmt.Println("unknown element inside map:", n.XMLName.Local)
		}
	}
}

func walkTMXLayer(l *pb.Layer, nodes []Node) {
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
			walkTMXData(d, n.Nodes)
			l.Data = d
		default:
			fmt.Println("unknown element inside layer:", n.XMLName.Local)
		}
	}
}

func walkTMXData(d *pb.Data, nodes []Node) {
	for _, n := range nodes {
		//Can contain: <tile>, <chunk>
		switch n.XMLName.Local {
		case "tile":
			t := &pb.DataTile{}
			for _, attr := range n.Attrs {
				switch attr.Name.Local {
				case "gid":
					t.Gid = toInt64(attr.Value)
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
					c.X = toInt64(attr.Value)
				case "y":
					c.Y = toInt64(attr.Value)
				case "width":
					c.Width = toInt64(attr.Value)
				case "height":
					c.Height = toInt64(attr.Value)
				default:
					fmt.Println("unknown attribute inside", n.XMLName.Local, "name:", attr.Name.Local, "value:", attr.Value)
				}
			}
			d.Chunks = append(d.Chunks, c)
		default:
			fmt.Println("unknown element inside layer:", n.XMLName.Local)
		}
	}
}
