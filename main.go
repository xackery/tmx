package main

import (
	"context"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/xackery/tmx/client"
	//"github.com/xackery/tmx/model"
)

var (
	version string
)

/*
gidBare := gid &^ GIDFlip
ID:             ID(gidBare - m.Tilesets[i].FirstGID)


770 //base val
771 //0 fff
2684355437 // 270 ccw tft
1073742701 // 180 ftf
536871683 // 90 ccw fft
val = val &^ (0x80000000 | 0x40000000 | 0x20000000)
*/
func main() {

	//fmt.Println(model.NewGID(2147483649).Index(), model.NewGID(2147483649).RotationRead())
	//return
	err := run()
	if err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}
}

func run() (err error) {

	var src, dst string
	args := os.Args
	if isVersionCheck(args) {
		return
	}
	if len(args) < 3 {
		usage()
		return
	}

	if len(args) > 2 {
		src = args[len(args)-2]
	}
	if len(args) > 1 {
		dst = args[len(args)-1]
	}

	if len(src) == 0 || len(dst) == 0 {
		usage()
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c, err := client.New(ctx)
	if err != nil {
		return
	}
	m, a, err := c.LoadFile(ctx, src)
	if err != nil {
		err = errors.Wrapf(err, "source %s", src)
		return
	}
	err = c.SaveFiles(ctx, m, a, dst)
	if err != nil {
		err = errors.Wrapf(err, "target %s", dst)
		return
	}
	return
}

func usage() (err error) {
	fmt.Println("usage: tmx source_file target_file")
	return
}

func isVersionCheck(args []string) (isCheck bool) {
	for _, arg := range args {
		if arg == "version" || arg == "/?" {
			fmt.Println("tmx version", version)
			isCheck = true
			return
		}
	}
	return
}
