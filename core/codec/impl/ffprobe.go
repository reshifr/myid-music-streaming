package codec_impl

import (
	"encoding/json"
	"strconv"
	"unicode"

	"github.com/reshifr/myid-music-streaming/core/codec"
	"github.com/reshifr/myid-music-streaming/core/ipc"
)

type FFprobe[CLI ipc.CLI] struct {
	cli CLI
}

func NewFFprobe[CLI ipc.CLI](cli CLI) FFprobe[CLI] {
	return FFprobe[CLI]{cli: cli}
}

func (ffprobe FFprobe[C]) Audio(path string) (*codec.AudioTag, error) {
	output, code := ffprobe.cli.Exec(
		"ffprobe",
		"-v", "-8",
		"-print_format", "json=c=1",
		"-show_entries", "format_tags=title,artist,album,genre,disc,track,date",
		path,
	)
	if code != ipc.ExitSuccess {
		return nil, codec.ErrTagReading
	}

	var data struct {
		Format struct {
			Tags struct {
				Title  string `json:"title"`
				Artist string `json:"artist"`
				Album  string `json:"album"`
				Genre  string `json:"genre"`
				Track  string `json:"track"`
				Disc   string `json:"disc"`
				Date   string `json:"date"`
			} `json:"tags"`
		} `json:"format"`
	}
	if err := json.Unmarshal(output, &data); err != nil {
		return nil, codec.ErrTagParsing
	}

	tags := data.Format.Tags
	tag := &codec.AudioTag{
		Title:  tags.Title,
		Artist: tags.Artist,
		Album:  tags.Album,
		Genre:  tags.Genre,
		Track:  parseTagNumber(tags.Track),
		Disc:   parseTagNumber(tags.Disc),
		Year:   parseTagNumber(tags.Date),
	}
	return tag, nil
}

func parseTagNumber(tagNumber string) (number uint32) {
	for i, r := range tagNumber {
		if !unicode.IsNumber(r) {
			tagNumber = tagNumber[:i]
			break
		}
	}
	parsedNumber, err := strconv.ParseUint(tagNumber, 10, 32)
	if err != nil {
		return 0
	}
	return uint32(parsedNumber)
}
