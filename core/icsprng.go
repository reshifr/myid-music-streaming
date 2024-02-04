package core

const (
	CSPRNGErrCode = 1
)

type ICSPRNG interface {
	Read(block []byte) (err error)
}
