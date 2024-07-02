package codec_test

import (
	"testing"

	codec_impl "github.com/reshifr/myid-music-streaming/core/codec/impl"
	ipc_impl "github.com/reshifr/myid-music-streaming/core/ipc/impl"
)

func Test_Playground(t *testing.T) {
	env := ipc_impl.Env{}
	cli := ipc_impl.NewStdCLI(env)
	tagReader := codec_impl.NewFFprobe(cli)
	tagReader.Audio("../../test_data/a.m4a")
}
