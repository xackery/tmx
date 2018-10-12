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

			//fmt.Println(newTs)

			ts = nt
		}

		if ts.Image == nil {
			fmt.Println("skipping empty image tilesheet")
			continue
		}
		sourceFile = ts.Image.Source
		imgPath := fmt.Sprintf("%s/%s", dir, sourceFile)
		a, err = NewAtlas(ctx, imgPath, nt)
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

	//extract image data

	ext := filepath.Ext(path)
	switch ext {
	case ".data":
		e := tmx.NewEncoder(f)
		err = e.Encode(m)
	case ".bin":
		e := tmx.NewEncoder(f)
		err = e.Encode(m)
	case ".xml":
		e := xml.NewEncoder(f)
		err = e.Encode(m)
	case ".yml":
		e := yaml.NewEncoder(f)
		err = e.Encode(m)
	case ".json":
		e := json.NewEncoder(f)
		err = e.Encode(m)
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
	//for now just png for atlas
	dir := filepath.Dir(path)
	atlasFile := "out.png"
	imgf, err := os.Create(fmt.Sprintf("%s/%s", dir, atlasFile))
	if err != nil {
		err = errors.Wrap(err, "failed to create file")
		return
	}
	defer imgf.Close()

	e := png.Encoder{
		CompressionLevel: png.BestCompression,
	}
	err = e.Encode(imgf, a.Image())
	if err != nil {
		err = errors.Wrapf(err, "encode %s/%s", dir, atlasFile)
		return
	}
	return
}
