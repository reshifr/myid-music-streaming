package crypt

import (
	"errors"
	"testing"

	"github.com/reshifr/play/core"
	core_mock "github.com/reshifr/play/mocks/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_ISO10126Pad_Add(t *testing.T) {
	t.Run("Length of block equals align", func(t *testing.T) {
		csprng := core_mock.NewCSPRNG(t)
		pad := InitISO10126Pad(csprng)
		align := uint8(16)
		bl := int(align)
		eblock := make([]byte, bl)
		var ecerr *core.Error = nil

		block, cerr := pad.Add(eblock, align)
		assert.Equal(t, eblock, block)
		assert.Equal(t, ecerr, cerr)
	})
	t.Run("CSPRNG read error", func(t *testing.T) {
		csprng := core_mock.NewCSPRNG(t)
		pad := InitISO10126Pad(csprng)
		align := uint8(16)
		bl := 20
		eblock := make([]byte, bl)
		var ecerr *core.Error = core.NewError(
			core.CSPRNGErrCode,
			"crypt.ISO10126Pad.Add(): rng read error.",
		)

		err := errors.New("read error")
		csprng.On("Read", mock.Anything).Return(err).Once()

		block, cerr := pad.Add(eblock, align)
		assert.Equal(t, eblock, block)
		assert.Equal(t, ecerr, cerr)
	})
	t.Run("Normal", func(t *testing.T) {
		csprng := core_mock.NewCSPRNG(t)
		pad := InitISO10126Pad(csprng)
		align := uint8(16)
		bl := 20
		eblock := make([]byte, bl)
		var ecerr *core.Error = nil

		zpadd := 12
		padd := make([]byte, zpadd)
		var err error = nil
		csprng.On("Read", padd).Return(err).Once()

		block, cerr := pad.Add(eblock, align)
		padd[zpadd-1] = byte(zpadd)
		eblock = append(eblock, padd...)
		assert.Equal(t, eblock, block)
		assert.Equal(t, ecerr, cerr)
	})
}
