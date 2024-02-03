package codec

import (
	"bytes"
	"strconv"

	"github.com/reshifr/play/core"
	"github.com/reshifr/play/core/ipc"
)

type FFmpeg[C ipc.ICLI] struct {
	cli C
}

func OpenFFmpeg[C ipc.ICLI](cli C) FFmpeg[C] {
	return FFmpeg[C]{cli: cli}
}

func (ffmpeg *FFmpeg[C]) ReadTag(path string) (*AudioTag, *core.Error) {
	output, cerr := ffmpeg.cli.Exec(
		"ffprobe",
		"-v", "-8",
		"-print_format", "flat",
		"-show_entries", "format_tags=title,artist,album,genre,disc,track",
		path,
	)
	if cerr != nil {
		return nil, cerr
	}
	flat := decodeFlat(output)
	tag := &AudioTag{
		Title:  flat["title"],
		Artist: flat["artist"],
		Album:  flat["album"],
		Genre:  flat["genre"],
		Track:  decodeOrder[uint16](flat["track"]),
		Disc:   decodeOrder[uint8](flat["disc"]),
	}
	return tag, cerr
}

func decodeFlat(data []byte) map[string]string {
	var vo bool = true
	var ika, ikb, iva, ivb int
	flat := make(map[string]string)
	for i := 0; i < len(data); i++ {
		switch data[i] {
		case byte('.'):
			ika = i + 1
		case byte('='):
			ikb = i
		case byte('"'):
			if vo {
				iva = i + 1
			} else {
				ivb = i
			}
			vo = !vo
		case byte('\n'):
			if ika <= ikb && iva <= ivb {
				k := string(bytes.ToLower(data[ika:ikb]))
				v := string(data[iva:ivb])
				flat[k] = v
			}
		}
	}
	return flat
}

func decodeOrder[O ~uint8 | ~uint16 | ~uint32](data string) O {
	var ie int
	for i, c := range data {
		if i == 0 && c == '0' {
			return 0
		}
		ie++
		if c < '0' || c > '9' {
			ie--
			break
		}
	}
	so := data[0:ie]
	o, err := strconv.ParseUint(so, 10, 64)
	if o > uint64(^O(0)) || err != nil {
		return 0
	}
	return O(o)
}
