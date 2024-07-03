package os

type File interface {
	Read(b []byte) (n int, err error)
	Seek(offset int64, whence int) (ret int64, err error)
	Close() error
}

type OS interface {
	Open(name string) (File, error)
}
