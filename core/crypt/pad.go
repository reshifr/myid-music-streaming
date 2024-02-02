package crypt

const (
	PAD_RNG_ERROR = 1
)

type Pad interface {
	Add(block []byte, aligned uint8) (padded []byte, ok bool)
}
