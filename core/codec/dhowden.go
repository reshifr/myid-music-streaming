package codec

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/reshifr/myid-music-streaming/common/drodhowden"
	"github.com/reshifr/myid-music-streaming/common/os"
)

type Drodhowden[OS os.OS, Tag drodhowden.Tag] struct {
	os  OS
	tag Tag
}

func NewDrodhowden[OS os.OS, Tag drodhowden.Tag](os OS, tag Tag) Drodhowden[OS, Tag] {
	return Drodhowden[OS, Tag]{os, tag}
}

func (Drodhowden[OS, Tag]) transcodeToJPEG(pict drodhowden.Picture) []byte {
	var err error
	var img image.Image
	var buf bytes.Buffer

	r := bytes.NewReader(pict.Data())
	img, err = gif.Decode(r)
	if err != nil {
		r.Seek(0, io.SeekStart)
		img, err = jpeg.Decode(r)
	}
	if err != nil {
		r.Seek(0, io.SeekStart)
		img, err = png.Decode(r)
	}
	if err != nil {
		return nil
	}

	err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: 100})
	if err != nil {
		return nil
	}
	return buf.Bytes()
}

func (d Drodhowden[OS, Tag]) AudioTag(path string) (*AudioTag, error) {
	f, err := d.os.Open(path)
	if err != nil {
		return nil, ErrTagReaderReadFile
	}
	defer func() {
		if fErr := f.Close(); fErr != nil {
			err = ErrTagReaderCloseFile
		}
	}()

	m, err := d.tag.ReadFrom(f)
	if err != nil {
		return nil, ErrTagReaderReadMetadata
	}

	track, _ := m.Track()
	disc, _ := m.Disc()
	tag := &AudioTag{
		Title:  m.Title(),
		Artist: m.Artist(),
		Album:  m.Album(),
		Cover:  d.transcodeToJPEG(m.Picture()),
		Track:  uint32(track),
		Disc:   uint32(disc),
	}
	return tag, err
}
