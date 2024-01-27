package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/hashicorp/golang-lru/v2"
	"github.com/reshifr/play/core/cli"
	"github.com/reshifr/play/core/codec"
)

func main() {
	cache, err := lru.New[string, string](10)
	if err != nil {
		fmt.Println("cache error!")
		os.Exit(1)
	}
	handler := cli.OSCmdHandler{
		Command: func(bin string, args ...string) cli.OSCmd {
			return exec.Command(bin, args...)
		},
		Clearenv: os.Clearenv,
	}
	ffmpeg := codec.OpenFFmpeg(cache, handler)
	ffmpeg.Nothing()
	// music, ok := codec.GetFlac("build/sia.flac")
	// if !ok {
	// 	fmt.Println("ERROR!")
	// 	return
	// }
	// fmt.Println(*music)

	// cache, ok := lru.New[string, string](100)
	// var a exec.Cmd
	// a.Output()
}
