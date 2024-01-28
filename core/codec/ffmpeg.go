package codec

import (
	"bytes"
	"strconv"

	"github.com/reshifr/play/core"
	"github.com/reshifr/play/core/ipc"
)

type FFmpeg struct {
	cli ipc.ICLI
}

func OpenFFmpeg(cli ipc.ICLI) (ffmpeg *FFmpeg) {
	ffmpeg = &FFmpeg{cli: cli}
	return ffmpeg
}

func (ffmpeg *FFmpeg) GetTag(path string) (tag *MusicTag, ok bool) {
	output, code := ffmpeg.cli.Exec(
		"ffprobe",
		"-v", "-8",
		"-print_format", "flat",
		"-show_entries", "format_tags=title,artist,album,genre,disc,track",
		path,
	)
	if code != core.CMD_EXIT_SUCCESS {
		return nil, false
	}
	flat := decodeFlat(output)
	tag = &MusicTag{
		Title:  flat["title"],
		Artist: flat["artist"],
		Album:  flat["album"],
		Genre:  flat["genre"],
		Disc:   extractOrder[uint32](flat["disc"]),
		Track:  extractOrder[uint32](flat["track"]),
	}
	return tag, true
}

func decodeFlat(data []byte) (flat map[string]string) {
	var vo bool = true
	var ika, ikb, iva, ivb int
	flat = make(map[string]string)
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

func extractOrder[O uint8 | uint16 | uint32 | uint64](data string) (order O) {
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
	o, _ := strconv.ParseUint(so, 10, 64)
	order = O(o)
	return order
}
