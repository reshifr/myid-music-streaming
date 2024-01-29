package codec

type IFFmpeg interface {
	ReadTag(path string) (tag *AudioTag, ok bool)
}
