package ipc

import (
	"github.com/reshifr/play/core"
)

const (
	CLIExitOk = iota
)

type ICLI interface {
	Exec(bin string, args ...string) (output []byte, cerr *core.Error)
}
