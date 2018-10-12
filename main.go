package main

import (
	"context"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/xackery/tmx/client"
	"github.com/xackery/tmx/model"
)

var (
	version string
)

func main() {
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
	if isVerboseCheck(args) {
		model.SetVerbose(true)
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

func isVerboseCheck(args []string) (isVerbose bool) {
	for _, arg := range args {
		if arg == "-v" {
			isVerbose = true
			return
		}
	}
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
