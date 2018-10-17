package atlas

import (
	"fmt"
	"image"
	"math"

	"github.com/xackery/tmx/model"
	"github.com/xackery/tmx/pb"
	"golang.org/x/image/draw"
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

// Bake returns the internal image
func (a *Atlas) Bake(m *pb.Map) (img *image.RGBA, err error) {

	//now we remap the map data with the new tilmap
	for i := range m.Layers {
		for j := 0; j < len(m.Layers[i].Data.DataTiles); j++ {
			d := m.Layers[i].Data.DataTiles[j]
			if d.Gid == 0 {
				continue
			}
			oldGid := model.NewGID(d.Gid)

			newGid := a.tileMap[model.Index(d.Gid)]
			if oldGid.HRead() && newGid.HRead() {
				newGid.HUpdate(false)
			}
			newGid.HUpdate(oldGid.HRead())
			if oldGid.VRead() && newGid.VRead() {
				newGid.VUpdate(false)
			}
			newGid.VUpdate(oldGid.VRead())
			if oldGid.DRead() && newGid.DRead() {
				newGid.DUpdate(false)
			}
			newGid.DUpdate(oldGid.DRead())

			fmt.Println(i, oldGid.Index(), oldGid.HRead(), oldGid.VRead(), oldGid.DRead(), oldGid.ValueRead(), newGid.ValueRead(), newGid.Index())
			//err = newGid.RotationUpdate(oldGid.RotationRead())
			//if err != nil {
			//	return
			//}
			//fmt.Println("remapped", oldGid.ValueRead(), oldGid.Index(), oldGid.RotationRead(), "to", newGid.ValueRead(), newGid.Index(), "rotation", newGid.RotationRead())
			if newGid == nil {
				continue
			}
			//newGid.RotationUpdate(oldGid.RotationRead())
			m.Layers[i].Data.DataTiles[j].Gid = newGid.ValueRead()
		}
	}

	//need to remove hard code
	//sheetWidth := 1024
	sheetWidth := 512
	img = image.NewRGBA(image.Rect(0, 0, sheetWidth, sheetWidth))
	//may not be needed
	//a.tileWidth += 60
	//a.tileHeight += 60
	tileSize := image.Rect(0, 0, int(a.tileWidth), int(a.tileHeight))
	offset := image.ZP
	//fmt.Println(len(a.tiles), "total tiles")
	for i := 0; i < len(a.tiles); i++ {
		tile := a.tiles[i]

		if tile == nil {
			err = fmt.Errorf("tile nil malfunction")
			return
		}
		//fmt.Printf("%d,%d|", offset.X, offset.Y)

		draw.Copy(img, offset, tile, tileSize, draw.Src, nil)
		offset.X += int(a.tileWidth)
		if offset.X > sheetWidth {
			offset.Y += int(a.tileHeight)
			offset.X = 0
		}
		if offset.Y > sheetWidth {
			fmt.Println("exceeded size")
			return
		}
	}
	a.img = img

	tileset := "tileset"
	if len(m.Layers) > 0 {
		tileset = "tilesets"
	}
	fmt.Println(tileset, "reduced from", a.oldTileCount, "to", a.newTileCount, "total tiles")
	return
}

func fastCompare(img1 *image.RGBA, img2 *image.NRGBA) (int64, error) {
	if img1.Bounds() != img2.Bounds() {
		return 0, fmt.Errorf("image bounds not equal: %+v, %+v", img1.Bounds(), img2.Bounds())
	}

	accumError := int64(0)

	for i := 0; i < len(img1.Pix); i++ {
		accumError += int64(sqDiffUInt8(img1.Pix[i], img2.Pix[i]))
	}

	return int64(math.Sqrt(float64(accumError))), nil
}

func sqDiffUInt8(x, y uint8) uint64 {
	d := uint64(x) - uint64(y)
	return d * d
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
	//var tx, ty int

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
			/*
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
			*/
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
	return
}
