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

// AppendUniqueThread will append if unique, otherwise returns existing index
func (a *Atlas) AppendUniqueThread(img *image.RGBA) (index int64) {

	var wg sync.WaitGroup
	matchChan := make(chan *scanImageResponse)
	doneChan := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rotation := 0
	//nImg := imaging.Rotate90(img)
	var tile *image.RGBA
	//doRotate := false
	for index = range a.tiles {
		rotation = 1
		tile = a.tiles[index]
		//tile = imaging.Rotate90(a.tiles[index])
		req := &scanImageRequest{
			ctx:       ctx,
			tile:      &image.RGBA{},
			img:       &image.RGBA{},
			wg:        &wg,
			matchChan: matchChan,
			index:     index,
			rotation:  rotation,
		}
		*req.tile = *tile
		*req.img = *img
		wg.Add(1)
		go scanImage(req)
		/*if doRotate {
			for i := 0; i < 3; i++ {
				rotation++
				//tile = imaging.Rotate90(tile)
				req := &scanImageRequest{
					ctx:       ctx,
					tile:      &image.RGBA{},
					img:       &image.RGBA{},
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
		*/
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
