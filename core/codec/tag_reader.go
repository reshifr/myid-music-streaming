package codec

type TagReaderError int

const (
	ErrTagReading TagReaderError = iota + 1
	ErrTagParsing
)

func (err TagReaderError) Error() string {
	switch err {
	case ErrTagReading:
		return "ErrInvalidIVLen: invalid IV length."
	default:
		return "Error: unknown."
	}
}

type AudioTag struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Album  string `json:"album"`
	Genre  string `json:"genre"`
	Track  uint32 `json:"track"`
	Disc   uint32 `json:"disc"`
	Year   uint32 `json:"year"`
}

type TagReader interface {
	Audio(path string) (tag *AudioTag, err error)
}
