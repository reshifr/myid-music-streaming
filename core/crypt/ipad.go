package crypt

import (
	"github.com/reshifr/play/core"
)

const (
	PadNullBlockError = iota + 1
	PadEmptyBlockError
	PadInvalidAlignError
)

type IPad interface {
	Add(block []byte, align uint8) (padded []byte, cerr *core.Error)
	Del(padded []byte, align uint8) (block []byte, cerr *core.Error)
}
