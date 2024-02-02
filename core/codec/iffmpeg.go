package codec

import (
	"github.com/reshifr/play/core"
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
	ReadTag(path string) (tag *AudioTag, err *core.Error)
}
