package codec_test

import (
	"encoding/json"
	"fmt"
	"testing"

	codec_impl "github.com/reshifr/myid-music-streaming/core/codec/impl"
	ipc_impl "github.com/reshifr/myid-music-streaming/core/ipc/impl"
)

func Test_Playground(t *testing.T) {
	env := ipc_impl.Env{}
	cli := ipc_impl.NewStdCLI(env)
	tagReader := codec_impl.NewFFprobe(cli)
	tag, _ := tagReader.Audio("../../../test_data/a.m4a")
	output, _ := json.MarshalIndent(tag, "", "  ")
	fmt.Println(string(output))
}
