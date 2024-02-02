package crypt

import (
	"github.com/reshifr/play/core"
)

type ISO10126Pad[Rng core.IRng] struct {
	rng Rng
}

func NewISO10126Pad[Rng core.IRng](rng Rng) *ISO10126Pad[Rng] {
	return &ISO10126Pad[Rng]{rng: rng}
}

func (pad *ISO10126Pad[Rng]) Add(
	block []byte,
	aligned uint8,
) ([]byte, *core.Error) {
	bl := uint(len(block))
	var coreErr *core.Error = nil
	if bl%uint(aligned) == 0 {
		return block, coreErr
	}
	bn := bl / uint(aligned)
	zpadd := (bn+1)*uint(aligned) - bl
	padd := make([]byte, zpadd)
	err := pad.rng.Read(padd)
	if err != nil {
		coreErr = core.ThrowError(
			PAD_RNG_ERROR,
			"crypt.ISO10126Pad.Add(): rng read error.",
		)
		return block, coreErr
	}
	padd[zpadd-1] = byte(zpadd)
	padded := append(block, padd...)
	return padded, coreErr
}
