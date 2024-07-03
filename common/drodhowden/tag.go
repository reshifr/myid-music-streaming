package drodhowden

import "io"

type Picture interface {
	Data() []byte
}

type Metadata interface {
	Title() string
	Album() string
	Artist() string
	Track() (int, int)
	Disc() (int, int)
	Picture() Picture
}

type Tag interface {
	ReadFrom(r io.ReadSeeker) (Metadata, error)
}
