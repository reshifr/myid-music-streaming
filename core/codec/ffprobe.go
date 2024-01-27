package codec

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/reshifr/play/core/cli"
	"github.com/reshifr/play/core/entity"
)

func GetFlac(path string) (music *entity.Music, ok bool) {
	c, ok := cli.Open()
	if !ok {
		log.Fatalf("Codec: can not open cli.")
		return nil, false
	}

	// tagTitle, titleStatus := cli.Exec(name, "--show-tag=TITLE", path)
	// tagArtist, artistStatus := cli.Exec(name, "--show-tag=ARTIST", path)
	// tagAlbum, albumStatus := cli.Exec(name, "--show-tag=ALBUM", path)
	// tagGenre, genreStatus := cli.Exec(name, "--show-tag=GENRE", path)
	music = &entity.Music{}
	ok = setTags(c, path, music)
	if !ok {
		return nil, false
	}
	// music.Title = getTag(tagTitle, titleStatus)
	// music.Artist = getTag(tagArtist, artistStatus)
	// music.Album = getTag(tagAlbum, albumStatus)
	// music.Genre = getTag(tagGenre, genreStatus)
	return music, true
}

func setTags(c *cli.Cli, path string, music *entity.Music) (ok bool) {
	output, code := c.Exec("ffprobe", "-v", "quiet",
		"-print_format", "json", "-show_entries", "format_tags", path)
	if code != cli.SUCCESS {
		return false
	}

	var data map[string]interface{}
	if err := json.Unmarshal(output, &data); err != nil {
		log.Fatalf("Codec: can not decode music tags.")
		return false
	}

	fmt.Println(data["format"]["tags"])

	return false
}

func getTag(output string, code int) string {
	if code != cli.SUCCESS {
		return ""
	}
	parts := strings.Split(output, "=")
	if len(parts) <= 1 {
		return ""
	}
	tag := parts[1]
	return tag
}
