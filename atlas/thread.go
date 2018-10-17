package atlas

import (
	"context"
	"image"
	"sync"
)

type scanImageRequest struct {
	ctx       context.Context
	tile      *image.RGBA
	img       *image.RGBA
	wg        *sync.WaitGroup
	matchChan chan *scanImageResponse
	index     int64
	rotation  int
}

type scanImageResponse struct {
	index    int64
	rotation int
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
