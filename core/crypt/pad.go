package crypt

import "github.com/reshifr/play/core"

const (
	PAD_RNG_ERROR = 1
)

type Pad[Rng core.IRng] interface {
	Add(block []byte, aligned uint8) (padded []byte, ok bool)
}
