package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	lru "github.com/hashicorp/golang-lru/v2"
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
	tag, _ := ffmpeg.GetTag("/home/reshifr/Downloads/sia")
	output, _ := json.MarshalIndent(tag, "", "  ")
	fmt.Println(string(output))
}
