package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/reshifr/play/core"
	"github.com/reshifr/play/core/codec"
)

func main() {
	handler := core.OSHandler{
		Clearenv: os.Clearenv,
		Command: func(path string, args ...string) (cmd core.OSCmd) {
			return exec.Command(path, args...)
		},
	}
	ffmpeg := codec.OpenFFmpeg(handler)
	tag, _ := ffmpeg.GetTag("/home/reshifr/Downloads/sia")
	output, _ := json.MarshalIndent(tag, "", "  ")
	fmt.Println(string(output))
}
