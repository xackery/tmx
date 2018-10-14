package client

import (
	"context"
	"os"

	"github.com/pkg/errors"
	"github.com/xackery/tmx/atlas"
	"github.com/xackery/tmx/pb"
)

// AppendAtlas appends elements to an atlas from an image file, using provided tileset as a reference
func AppendAtlas(ctx context.Context, a *atlas.Atlas, path string, m *pb.Map, t *pb.TileSet, usedTiles map[uint32]bool) (na *atlas.Atlas, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	d := atlas.NewDecoder(f, m, t, usedTiles)
	err = d.Decode(a)
	if err != nil {
		err = errors.Wrap(err, "marshal")
		return
	}
	na = a
	return
}
