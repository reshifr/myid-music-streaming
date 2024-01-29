package main

import (
	"encoding/json"
	"fmt"

	"github.com/reshifr/play/core"
	"github.com/reshifr/play/core/codec"
	"github.com/reshifr/play/core/ipc"
)

func main() {
	env := core.Env{}
	cli := ipc.OpenCLI(env)
	ffmpeg := codec.OpenFFmpeg(cli)
	tag, _ := ffmpeg.ReadTag("/home/reshifr/Downloads/sia")
	output, _ := json.MarshalIndent(tag, "", "  ")
	fmt.Println(string(output))
}
