package codec

type AudioTag struct {
	Title  string
	Artist string
	Album  string
	Genre  string
	Track  uint32
	Disc   uint16
}

type TagReader interface {
	Audio(path string) (tag *AudioTag, err error)
}
