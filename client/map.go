package client

import (
	"context"
	"os"

	"github.com/pkg/errors"
	"github.com/xackery/tmx/pb"
	"github.com/xackery/tmx/tmx"
)

// NewMap creates a new map (*.tmx), using provided path
func NewMap(ctx context.Context, path string) (m *pb.Map, err error) {
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
		return
	}
	return
}
