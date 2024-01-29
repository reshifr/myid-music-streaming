package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/reshifr/play/core"
	"github.com/reshifr/play/core/codec"
	"github.com/reshifr/play/core/ipc"
)

func main() {
	handler := core.HOS{
		Clearenv: os.Clearenv,
		Command: func(path string, args ...string) (cmd core.IOSCmd) {
			return exec.Command(path, args...)
		},
	}
	cli := ipc.OpenCLI(handler)
	ffmpeg := codec.OpenFFmpeg(cli)
	tag, _ := ffmpeg.ReadTag("/home/reshifr/Downloads/sia")
	output, _ := json.MarshalIndent(tag, "", "  ")
	fmt.Println(string(output))
}
