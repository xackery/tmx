package client

import (
	"context"
	"os"

	"github.com/pkg/errors"
	"github.com/xackery/tmx/atlas"
	"github.com/xackery/tmx/pb"
)

// NewAtlas creates a new atlas from an image file, using provided tileset as a reference
func NewAtlas(ctx context.Context, path string, t *pb.TileSet) (a *atlas.Atlas, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	d := atlas.NewDecoder(f, t)
	a = &atlas.Atlas{}
	err = d.Decode(a)
	if err != nil {
		err = errors.Wrap(err, "marshal")
		return
	}
	return
}
