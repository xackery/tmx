package atlas

import (
	"image"

	"github.com/xackery/tmx/model"
)

// Atlas represents an image atlas
type Atlas struct {
	tiles        map[int]*image.RGBA
	tileMap      map[uint32]*model.GID
	img          *image.RGBA
	tileWidth    int64
	tileHeight   int64
	newTileCount int
	oldTileCount int
}

// LastTileIndex returns the last added index
func (a *Atlas) LastTileIndex() int {
	return len(a.tiles)
}

// AppendUnique will append if unique and return new index, otherwise returns existing index
func (a *Atlas) AppendUnique(img *image.RGBA) (gid *model.GID) {
	var isMatch, h, v, d bool
	for i := 0; i < len(a.tiles); i++ {
		isMatch, h, v, d = doCompare(img, a.tiles[i])
		if isMatch {
			gid = model.NewGID(uint32(i + 1))
			gid.HUpdate(h)
			gid.VUpdate(v)
			gid.DUpdate(d)
			return
		}
	}
	a.tiles[len(a.tiles)] = img
	gid = model.NewGID(uint32(len(a.tiles)))
	return
}

func doCompare(img *image.RGBA, tile *image.RGBA) (isMatch, h, v, d bool) {
	isMatch = true //any orientation matched
	m := true      //normal match
	h = true       //horizontal flag match
	v = true       //vertical flag match
	d = true       //diagnol flag match
	hvd := true
	hd := true
	vd := true
	var tx, ty int

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			//m
			if m && tile.At(x, y) != img.At(x, y) {
				m = false
			}
			h = false
			v = false
			d = false
			hvd = false
			hd = false
			vd = false

			//h
			tx = img.Bounds().Max.X - x
			ty = y
			if h && tile.At(tx, ty) != img.At(x, y) {
				h = false
			}
			//v
			tx = x
			ty = img.Bounds().Max.Y - y
			if v && tile.At(tx, ty) != img.At(x, y) {
				v = false
			}
			//d
			tx = img.Bounds().Max.Y - x
			ty = x
			if d && tile.At(tx, ty) != img.At(x, y) {
				d = false
			}

			//hvd
			tx = img.Bounds().Max.Y - x
			ty = x
			tx = img.Bounds().Max.X - tx
			ty = img.Bounds().Max.Y - ty
			if hvd && tile.At(tx, ty) != img.At(x, y) {
				hvd = false
			}

			//vd
			tx = img.Bounds().Max.Y - x
			ty = x
			ty = img.Bounds().Max.Y - ty
			if vd && tile.At(tx, ty) != img.At(x, y) {
				vd = false
			}

			//hd
			tx = img.Bounds().Max.Y - x
			ty = x
			tx = img.Bounds().Max.X - tx
			if vd && tile.At(tx, ty) != img.At(x, y) {
				vd = false
			}

			if !m && !h && !v && !d && !hvd && !hd && !vd {
				isMatch = false
				return
			}
		}
	}
	if hvd {
		h = true
		v = true
		d = true
	}
	if hd {
		h = true
		d = true
	}
	if vd {
		v = true
		d = true
	}
	isMatch = true
	return
}
