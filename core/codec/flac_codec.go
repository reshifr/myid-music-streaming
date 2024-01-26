package codec

import (
	"github.com/reshifr/play/core/cli"
	"github.com/reshifr/play/core/entity"
	"strings"
)

func GetFlac(path string) *entity.Music {
	name := "metaflac"
	var cmd cli.Cli
	var music entity.Music
	tagTitle, titleStatus := cmd.Exec(name, "--show-tag=TITLE", path)
	tagArtist, artistStatus := cmd.Exec(name, "--show-tag=ARTIST", path)
	tagAlbum, albumStatus := cmd.Exec(name, "--show-tag=ALBUM", path)
	tagGenre, genreStatus := cmd.Exec(name, "--show-tag=GENRE", path)
	music.Title = getCleanTag(tagTitle, titleStatus)
	music.Artist = getCleanTag(tagArtist, artistStatus)
	music.Album = getCleanTag(tagAlbum, albumStatus)
	music.Genre = getCleanTag(tagGenre, genreStatus)
	return &music
}

func getCleanTag(tag string, status int) string {
	if status == cli.FAILURE {
		return ""
	}
	subTags := strings.Split(tag, "=")
	if len(subTags) <= 1 {
		return ""
	}
	return subTags[1]
}
