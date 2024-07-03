package codec

type TagReaderError int

const (
	ErrTagReaderReadFile TagReaderError = iota + 1
	ErrTagReaderReadMetadata
	ErrTagReaderCloseFile
)

func (err TagReaderError) Error() string {
	switch err {
	case ErrTagReaderReadFile:
		return "ErrTagReaderReadFile: hello."
	default:
		return "Error: unknown."
	}
}

type AudioTag struct {
	Title  string
	Artist string
	Album  string
	Cover  []byte
	Track  uint32
	Disc   uint32
}

type TagReader interface {
	AudioTag(path string) (tag *AudioTag, err error)
}
