package atlas

import (
	"context"
	"fmt"
	"image"
	"math"
	"sync"

	"github.com/disintegration/imaging"
)

// Atlas represents an image atlas
type Atlas struct {
	tiles   map[int64]*image.RGBA
	tileMap map[int64]int64
	img     *image.RGBA
}

type scanImageRequest struct {
	ctx       context.Context
	tile      *image.NRGBA
	img       *image.NRGBA
	wg        *sync.WaitGroup
	matchChan chan *scanImageResponse
	index     int64
	rotation  int
}

type scanImageResponse struct {
	index    int64
	rotation int
}

// AppendUniqueThread will append if unique, otherwise returns existing index
func (a *Atlas) AppendUniqueThread(img *image.RGBA) (index int64) {
	var wg sync.WaitGroup
	matchChan := make(chan *scanImageResponse)
	doneChan := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rotation := 0
	nImg := imaging.Rotate90(img)
	var tile *image.NRGBA
	for index = range a.tiles {
		rotation = 1
		tile = imaging.Rotate90(a.tiles[index])
		req := &scanImageRequest{
			ctx:       ctx,
			tile:      &image.NRGBA{},
			img:       &image.NRGBA{},
			wg:        &wg,
			matchChan: matchChan,
			index:     index,
			rotation:  rotation,
		}
		*req.tile = *tile
		*req.img = *nImg
		wg.Add(1)
		go scanImage(req)
		for i := 0; i < 3; i++ {
			rotation++
			tile = imaging.Rotate90(tile)
			req := &scanImageRequest{
				ctx:       ctx,
				tile:      &image.NRGBA{},
				img:       &image.NRGBA{},
				wg:        &wg,
				matchChan: matchChan,
				index:     index,
				rotation:  rotation,
			}
			*req.tile = *tile
			*req.img = *nImg
			wg.Add(1)
			go scanImage(req)
		}

	}
	go func() {
		wg.Wait()
		doneChan <- true
	}()

	select {
	case resp := <-matchChan:
		index = resp.index
		//fmt.Println("winner:", resp.rotation, resp.index)
	case <-doneChan:
		//fmt.Println("make new")
		index = int64(len(a.tiles))
		a.tiles[index] = img
	}
	return
}

func scanImage(req *scanImageRequest) {
	img := req.img
	tile := req.tile
	defer req.wg.Done()
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			select {
			case <-req.ctx.Done():
				return
			default:
			}
			if tile.At(x, y) != img.At(x, y) {
				//fmt.Println("not", req.rotation, req.index)
				return
			}
		}
	}
	req.matchChan <- &scanImageResponse{index: req.index, rotation: req.rotation}
}

// Image returns the internal image
func (a *Atlas) Image() (img *image.RGBA) {
	img = a.img
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

// AppendUnique will append if unique, otherwise returns existing index
func (a *Atlas) AppendUnique(img *image.RGBA) (index int64) {
	var tile *image.RGBA
	var i *image.NRGBA
	var t *image.NRGBA
	for index, tile = range a.tiles {
		i = imaging.Rotate90(img)
		t = imaging.Rotate90(tile)
		isMatch := doCompare(i, t)
		if isMatch {
			return
		}
		i = imaging.Rotate90(img)
		t = imaging.Rotate90(tile)
		isMatch = doCompare(i, t)
		if isMatch {
			return
		}
		i = imaging.Rotate90(img)
		t = imaging.Rotate90(tile)
		isMatch = doCompare(i, t)
		if isMatch {
			return
		}
		i = imaging.Rotate90(img)
		t = imaging.Rotate90(tile)
		isMatch = doCompare(i, t)
		if isMatch {
			return
		}
	}

	index = int64(len(a.tiles))
	a.tiles[index] = img
	return
}

func doCompare(img *image.NRGBA, tile *image.NRGBA) (isMatch bool) {
	isMatch = true
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
	return
}
