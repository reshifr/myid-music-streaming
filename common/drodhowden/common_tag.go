package drodhowden

import (
	"io"

	"github.com/dhowden/tag"
)

type commonPicture struct {
	*tag.Picture
}

func (c commonPicture) Data() []byte { return c.Picture.Data }

type commonMetadata struct {
	tag.Metadata
}

func (c commonMetadata) Picture() Picture { return commonPicture{c.Metadata.Picture()} }

type CommonTag struct{}

func (CommonTag) ReadFrom(r io.ReadSeeker) (Metadata, error) {
	m, err := tag.ReadFrom(r)
	return commonMetadata{m}, err
}
