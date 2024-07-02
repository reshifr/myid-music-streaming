package ipc_impl

import (
	"os/exec"

	"github.com/reshifr/myid-music-streaming/core/ipc"
)

type Env struct{}

func (env Env) Command(bin string, args ...string) ipc.Cmd {
	args = append([]string{bin}, args...)
	return exec.Command("/usr/bin/env", args...)
}
