package core

const (
	CSPRNGReadErr = iota + 1
)

type ICSPRNG interface {
	Read(block []byte) (err error)
}
