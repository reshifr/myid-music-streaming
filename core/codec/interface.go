package codec

type IFFmpeg interface {
	GetTag(path string) (tag *MusicTag, ok bool)
}
