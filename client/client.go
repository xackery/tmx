package client

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
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

// LoadFile loads a TMX file
func (c *Client) LoadFile(ctx context.Context, path string) (m *pb.Map, err error) {
	f, err := os.Open(path)
	if err != nil {
		err = errors.Wrap(err, "failed to load")
		return
	}
	defer f.Close()
	m = &pb.Map{}
	d := tmx.NewDecoder(f)
	err = d.Decode(m)
	if err != nil {
		err = errors.Wrap(err, "failed to marshal")
	}
	return
}

// SaveFile saves protobuf binary data
func (c *Client) SaveFile(ctx context.Context, m *pb.Map, path string) (err error) {
	f, err := os.Create(path)
	if err != nil {
		err = errors.Wrap(err, "failed to create file")
		return
	}
	defer f.Close()

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

	return
}
