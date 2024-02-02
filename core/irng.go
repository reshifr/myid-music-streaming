package core

const (
	RNGErrCode = 1
)

type IRNG interface {
	Read(block []byte) (err error)
}
