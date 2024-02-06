package codec

import (
	"github.com/reshifr/play/core"
)

const (
	FFmpegExecErr = iota + 1
)

type AudioTag struct {
	Title  string
	Artist string
	Album  string
	Genre  string
	Track  uint16
	Disc   uint8
}

type IFFmpeg interface {
	ReadTag(path string) (tag *AudioTag, cerr *core.Error)
}
