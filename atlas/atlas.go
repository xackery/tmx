package atlas

import (
	"image"
)

// Atlas represents an image atlas
type Atlas struct {
	tiles   map[int64]*image.RGBA
	tileMap map[int64]int64
	img     *image.RGBA
}

// AppendUnique will append if unique, otherwise returns existing index
func (a *Atlas) AppendUnique(img *image.RGBA) (index int64) {
	var tile *image.RGBA
	for index, tile = range a.tiles {
		isMatch := true
		for y := img.Bounds().Min.Y; isMatch && y < img.Bounds().Max.Y; y++ {
			for x := img.Bounds().Min.X; isMatch && x < img.Bounds().Max.X; x++ {
				//fmt.Println(tile.At(x, y), img.At(x, y))
				if tile.At(x, y) != img.At(x, y) {
					isMatch = false
				}
			}
		}
		if isMatch {
			return
		}
	}

	index = int64(len(a.tiles))
	a.tiles[index] = img
	return
}

// Image returns the internal image
func (a *Atlas) Image() (img *image.RGBA) {
	img = a.img
	return
}
