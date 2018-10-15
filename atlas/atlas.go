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
	tiles        map[int64]*image.RGBA
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
		for j, d := range m.Layers[i].Data.DataTiles {
			if d.Gid == 0 {
				continue
			}
			oldGid := model.NewGID(d.Gid)

			newGid := a.tileMap[model.Index(d.Gid)]
			err = newGid.RotationUpdate(oldGid.RotationRead())
			if err != nil {
				return
			}
			//fmt.Println("remapped", oldGid.ValueRead(), oldGid.Index(), oldGid.RotationRead(), "to", newGid.ValueRead(), newGid.Index(), "rotation", newGid.RotationRead())
			if newGid == nil {
				continue
			}
			newGid.RotationUpdate(oldGid.RotationRead())
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
	for _, tile := range a.tiles {

		offset.X += int(a.tileWidth)
		if offset.X > sheetWidth {
			offset.Y += int(a.tileHeight)
			offset.X = 0
		}
		if offset.Y > sheetWidth {
			fmt.Println("exceeded size")
			return
		}
		if tile == nil {
			err = fmt.Errorf("tile nil malfunction")
			return
		}
		//fmt.Printf("%d,%d|", offset.X, offset.Y)

		draw.Copy(img, offset, tile, tileSize, draw.Src, nil)
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
	var index int64
	var rotation int
	defer func() {
		gid = model.NewGID(uint32(index))
		switch rotation {
		case 0: //no rotation
		case 1: //90 ccw
			gid.RotationUpdate(270)
		case 2:
			gid.RotationUpdate(180)
		case 3:
			gid.RotationUpdate(90)
		}
	}()

	var tile *image.RGBA
	for index, tile = range a.tiles {
		for rotation = 0; rotation < 4; rotation++ {
			if doCompare(img, tile, rotation) {
				return
			}
		}
	}
	a.tiles[int64(len(a.tiles))] = img
	index = int64(len(a.tiles))
	return
}

func doCompare(img *image.RGBA, tile *image.RGBA, rotation int) (isMatch bool) {
	isMatch = true
	switch rotation {
	case 0: //0*
		for y := img.Bounds().Min.Y; isMatch && y < img.Bounds().Max.Y; y++ {
			for x := img.Bounds().Min.X; isMatch && x < img.Bounds().Max.X; x++ {
				//fmt.Println(tile.At(x, y), img.At(x, y))
				if tile.At(x, y) != img.At(x, y) {
					isMatch = false
					return
				}
			}
		}
	case 1: //90*

		// ix = image
		// x = tile

		ix := 0
		iy := 0

		//y = 0, y< max; y++
		//x = 0, x< max; x++
		//0,0 0,1 0,2
		//1,0 1,1 1,2
		//2,0 2,1 2,2
		//CC W90
		//x=2,y=1

		for y := img.Bounds().Min.Y; isMatch && y < img.Bounds().Max.Y; y++ {
			for x := img.Bounds().Min.X; isMatch && x < img.Bounds().Max.X; x++ {
				ix = y
				iy = x
				if tile.At(x, y) != img.At(ix, iy) {
					isMatch = false
					return
				}
			}
		}
	case 2: //180*
		//x0, y1
		//0,0 0,1 0,2
		//1,0 1,1 1,2
		//2,0 2,1 2,2
		//x = iy
		//y = ix
		ix := 0
		iy := 0

		for y := img.Bounds().Min.Y; isMatch && y < img.Bounds().Max.Y; y++ {
			for x := img.Bounds().Min.X; isMatch && x < img.Bounds().Max.X; x++ {
				ix = img.Bounds().Max.X - x
				iy = img.Bounds().Max.Y - y
				if tile.At(x, y) != img.At(ix, iy) {
					isMatch = false
					return
				}
			}
		}
	case 3: //270*
		ix := 0
		iy := 0
		//0,0 0,1 0,2, 0,3 .. 0,7, 0,8
		//1,0 1,1 1,2,             1,8
		//2,0 2,1 2,2
		//..
		//7,0
		//8,0

		// x = 2, y = 1
		// ix = (maxX - x)
		// ix = (2-1), = 1
		// iy = (maxY - y)
		// iy = (2-2) = 0

		// x = 7, y = 0
		// ix = 8-7 = 1
		// iy = 8-0= 8
		for y := img.Bounds().Min.Y; isMatch && y < img.Bounds().Max.Y; y++ {
			for x := img.Bounds().Min.X; isMatch && x < img.Bounds().Max.X; x++ {
				ix = img.Bounds().Max.X - x
				iy = img.Bounds().Max.Y - y
				if tile.At(x, y) != img.At(ix, iy) {
					isMatch = false
					return
				}

			}
		}
	}
	if isMatch {
		return
	}
	return
}
