package crypt

import (
	"testing"

	"github.com/reshifr/play/core"
	core_mock "github.com/reshifr/play/mocks/core"
	"github.com/stretchr/testify/assert"
)

func Test_ISO10126Pad_Add(t *testing.T) {
	t.Run("Len of block equals align", func(t *testing.T) {
		rng := core_mock.NewRNG(t)
		pad := InitISO10126Pad(rng)
		align := uint8(8)
		block := make([]byte, align)
		var coreErr *core.Error = nil
		expBlock, expCoreErr := pad.Add(block, align)
		assert.Equal(t, expBlock, block)
		assert.Equal(t, expCoreErr, coreErr)
	})
}
