package crypt

import (
	"github.com/reshifr/play/core"
)

type ISO10126Pad[R core.IRNG] struct {
	rng R
}

func InitISO10126Pad[R core.IRNG](rng R) ISO10126Pad[R] {
	return ISO10126Pad[R]{rng: rng}
}

func (pad *ISO10126Pad[R]) Add(block []byte, align uint8) ([]byte, *core.Error) {
	bl := uint(len(block))
	var cerr *core.Error = nil
	if bl%uint(align) == 0 {
		return block, cerr
	}
	bn := bl / uint(align)
	zpadd := (bn+1)*uint(align) - bl
	padd := make([]byte, zpadd)
	err := pad.rng.Read(padd)
	if err != nil {
		cerr = core.NewError(
			core.RNGErrCode,
			"crypt.ISO10126Pad.Add(): rng read error.",
		)
		return block, cerr
	}
	padd[zpadd-1] = byte(zpadd)
	padded := append(block, padd...)
	return padded, cerr
}
