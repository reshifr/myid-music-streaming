package codec_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/reshifr/myid-music-streaming/common/drodhowden"
	"github.com/reshifr/myid-music-streaming/common/os"
	"github.com/reshifr/myid-music-streaming/core/codec"
)

func Test_Playground(t *testing.T) {
	os := os.CommonOS{}
	tag := drodhowden.CommonTag{}
	tagReader := codec.NewDrodhowden(os, tag)
	audioTag, err := tagReader.AudioTag("../../test_data/c.mp3")
	fmt.Println(err)
	output, err := json.MarshalIndent(audioTag, "", "  ")
	fmt.Println(err)
	fmt.Println(string(output))
}
