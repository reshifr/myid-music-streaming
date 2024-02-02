package crypt

import (
	"github.com/reshifr/play/core"
)

type ISO10126Pad[RNG core.IRNG] struct {
	rng RNG
}

func InitISO10126Pad[RNG core.IRNG](rng RNG) *ISO10126Pad[RNG] {
	return &ISO10126Pad[RNG]{rng: rng}
}

func (pad *ISO10126Pad[RNG]) Add(
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
			core.RNGErrCode,
			"crypt.ISO10126Pad.Add(): rng read error.",
		)
		return block, coreErr
	}
	padd[zpadd-1] = byte(zpadd)
	padded := append(block, padd...)
	return padded, coreErr
}
