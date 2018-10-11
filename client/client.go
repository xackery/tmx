package client

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/pkg/errors"
	"github.com/xackery/tmx/pb"
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
	d := NewDecoder(f)
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
		e := NewEncoder(f)
		err = e.Encode(m)
	case ".bin":
		e := NewEncoder(f)
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

func toInt64(source string) (value int64) {
	value, err := strconv.ParseInt(source, 10, 64)
	if err != nil {
		fmt.Println("invalid integer value:", source)
	}
	return
}

func toBool(source string) (value bool) {
	val, err := strconv.ParseInt(source, 10, 64)
	if err != nil {
		fmt.Println("invalid integer value:", source)
	}
	value = val > 0
	return
}

func toFloat32(source string) (value float32) {
	val, err := strconv.ParseFloat(source, 32)
	if err != nil {
		fmt.Println("invalid float value:", source)
	}
	value = float32(val)
	return
}
