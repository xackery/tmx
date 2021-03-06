package client

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"image/png"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/xackery/tmx/atlas"
	"github.com/xackery/tmx/model"
	"github.com/xackery/tmx/pb"
	"github.com/xackery/tmx/tmx"
	"gopkg.in/yaml.v2"
)

// Client represents the TMX library
type Client struct {
}

// New creates a new TMX client library
func New(ctx context.Context) (c *Client, err error) {
	return
}

// LoadFile loads a TMX file and all referenced assets
func (c *Client) LoadFile(ctx context.Context, path string) (m *pb.Map, a *atlas.Atlas, err error) {
	dir := filepath.Dir(path)
	//load map
	m, err = NewMap(ctx, path)
	if err != nil {
		err = errors.Wrap(err, "failed to load tmx")
		return
	}

	usedTiles := make(map[uint32]bool)
	//make a list of every single tile used
	for _, l := range m.Layers {
		for _, d := range l.Data.DataTiles {
			usedTiles[model.Index(d.Gid)] = true
		}
	}

	a = &atlas.Atlas{}
	var nt *pb.TileSet
	//load external tilesets
	for i := range m.TileSets {
		ts := m.TileSets[i]

		firstGid := ts.FirstGid
		sourceFile := ts.Source
		if len(ts.Source) != 0 {
			nt, err = NewTileSet(ctx, fmt.Sprintf("%s/%s", dir, sourceFile))
			if err != nil {
				err = errors.Wrapf(err, "tsx %d %s", firstGid, sourceFile)
				return
			}

			nt.FirstGid = firstGid
			nt.Source = sourceFile
			m.TileSets[i] = nt
			ts = nt
		}

		if ts.Image == nil {
			fmt.Println("skipping empty image tilesheet")
			continue
		}
		sourceFile = ts.Image.Source
		imgPath := fmt.Sprintf("%s/%s", dir, sourceFile)
		fmt.Println("parsing tileset", imgPath)
		a, err = AppendAtlas(ctx, a, imgPath, m, nt, usedTiles)

		if err != nil {
			err = errors.Wrapf(err, "tsx %d %s (img: %s)", firstGid, sourceFile, imgPath)
			return
		}
	}
	return
}

// SaveFiles saves protobuf binary data and atlas image
func (c *Client) SaveFiles(ctx context.Context, m *pb.Map, a *atlas.Atlas, path string) (err error) {
	f, err := os.Create(path)
	if err != nil {
		err = errors.Wrap(err, "failed to create file")
		return
	}
	defer f.Close()
	ext := filepath.Ext(path)

	iPath := path[0 : len(path)-len(ext)]
	//for now just png for atlas
	//dir := filepath.Dir(path)
	atlasFile := ".png"
	imgPath := fmt.Sprintf("%s%s", iPath, atlasFile)
	imgf, err := os.Create(imgPath)
	if err != nil {
		err = errors.Wrap(err, "failed to create file")
		return
	}
	defer imgf.Close()

	e := png.Encoder{
		CompressionLevel: png.BestCompression,
	}
	img, err := a.Bake(m)
	if err != nil {
		err = errors.Wrap(err, "failed to encode new atlas image")
		return
	}
	err = e.Encode(imgf, img)
	if err != nil {
		err = errors.Wrapf(err, "atlasEncode %s", imgPath)
		return
	}

	//extract image data

	out := &pb.OutMap{
		Source:     imgPath,
		Width:      m.Width,
		Height:     m.Height,
		TileWidth:  m.TileWidth,
		TileHeight: m.TileHeight,
		TileCount:  int64(a.LastTileIndex()),
		Colliders:  m.Colliders,
	}
	for _, l := range m.Layers {
		outL := &pb.OutLayer{
			Name:    l.Name,
			Opacity: l.Opacity,
		}
		for _, t := range l.Data.DataTiles {
			outT := &pb.OutTile{Gid: t.GetGid()}
			outL.Tiles = append(outL.Tiles, outT)
		}
		out.Layers = append(out.Layers, outL)
	}

	//do layers
	switch ext {
	case ".data":
		e := tmx.NewEncoder(f)
		err = e.Encode(out)
	case ".bin":
		e := tmx.NewEncoder(f)
		err = e.Encode(out)
	case ".xml":
		e := xml.NewEncoder(f)
		err = e.Encode(out)
	case ".yml":
		e := yaml.NewEncoder(f)
		err = e.Encode(out)
	case ".json":
		e := json.NewEncoder(f)
		e.SetIndent("", "	")
		err = e.Encode(out)
	default:
		err = fmt.Errorf("unknown target file extension: %s (accepted values: json, data, bin)", ext)
		return
	}
	if err != nil {
		err = errors.Wrapf(err, "failed to encode to %s", ext)
		return
	}

	if a == nil {
		fmt.Println("warning: no atlas generated")
		return
	}

	return
}
