package codec

import (
	// "encoding/json"
	// "fmt"
	// "log"
	// "strings"

	"github.com/reshifr/play/core/cli"
	// "github.com/reshifr/play/core/entity"
	// "github.com/reshifr/play/core/entity"
)

type FFmpeg struct {
	cmd *cli.Cmd
}

func OpenFFmpeg(cache cli.LRU, handler cli.OSCmdHandler) (ffmpeg *FFmpeg) {
	cmd := cli.OpenCmd(cache, handler)
	ffmpeg = &FFmpeg{cmd: cmd}
	return ffmpeg
}

func (ffmpeg *FFmpeg) Nothing() {}

// func (ffmpeg *FFmpeg) GetMusic(path string) (music *entity.Music, ok bool) {
// 	music = &entity.Music{}
// 	ok = setTags(c, path, music)
// 	if !ok {
// 		return nil, false
// 	}
// 	return music, true
// }

// func setTags(c *cli.Cli, path string, music *entity.Music) (ok bool) {
// 	output, code := c.Exec(
// 		"ffprobe",
// 		"-v", "quiet",
// 		"-print_format", "json",
// 		"-show_entries", "format_tags",
// 		path
// 	)
// 	if code != cli.SUCCESS {
// 		return false
// 	}

// 	// var data map[string]interface{}
// 	// if err := json.Unmarshal(output, &data); err != nil {
// 	// 	log.Fatalf("Codec: can not decode music tags.")
// 	// 	return false
// 	// }

// 	// fmt.Println(data["format"]["tags"])

// 	return false
// }

// func getTag(output string, code int) string {
// 	if code != cli.SUCCESS {
// 		return ""
// 	}
// 	parts := strings.Split(output, "=")
// 	if len(parts) <= 1 {
// 		return ""
// 	}
// 	tag := parts[1]
// 	return tag
// }
