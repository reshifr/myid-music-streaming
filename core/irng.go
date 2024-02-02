package core

type IRng interface {
	Read(block []byte) (err error)
}
