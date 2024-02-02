package crypt

type IPad interface {
	Add(block []byte, aligned uint8) (padded []byte, ok bool)
}
