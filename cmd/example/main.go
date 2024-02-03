package main

import (
	"encoding/json"
	"fmt"

	"github.com/reshifr/play/core"
	"github.com/reshifr/play/core/codec"
	"github.com/reshifr/play/core/ipc"
)

func main() {
	var env core.Env
	cli := ipc.OpenCLI(&env)
	ffmpeg := codec.OpenFFmpeg(&cli)
	tag, cerr := ffmpeg.ReadTag("build/x1")
	output, _ := json.MarshalIndent(tag, "", "  ")
	if cerr != nil {
		fmt.Println(cerr.Msg)
		return
	}
	stag := string(output)
	fmt.Println(stag)
}
