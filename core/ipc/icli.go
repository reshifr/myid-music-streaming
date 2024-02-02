package ipc

import (
	"github.com/reshifr/play/core"
)

type ICLI interface {
	Exec(bin string, args ...string) (output []byte, err *core.Error)
}
