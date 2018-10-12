package client

import (
	"context"
	"os"

	"github.com/pkg/errors"
	"github.com/xackery/tmx/pb"
	"github.com/xackery/tmx/tsx"
)

// NewTileSet creates a new tileset (*.tsx), using provided path
func NewTileSet(ctx context.Context, path string) (t *pb.TileSet, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	t = &pb.TileSet{}
	d := tsx.NewDecoder(f)
	err = d.Decode(t)
	if err != nil {
		err = errors.Wrap(err, "marshal")
		return
	}
	return
}
