package main

import (
	"context"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/xackery/tmx/client"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func run() (err error) {
	var src, dst string
	args := os.Args
	if len(args) < 3 {
		usage()
		return
	}

	if len(args) > 1 {
		src = args[len(args)-2]
	}
	if len(args) > 2 {
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
	m, err := c.LoadFile(ctx, src)
	if err != nil {
		err = errors.Wrapf(err, "failed to load from %s", src)
		return
	}
	err = c.SaveFile(ctx, m, dst)
	if err != nil {
		err = errors.Wrapf(err, "failed to save to %s", dst)
		return
	}
	return
}

func usage() (err error) {
	fmt.Println("usage: tmx source_file target_file")
	return
}
