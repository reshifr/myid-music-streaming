package core

import (
	"os"
	"os/exec"
)

type Env struct{}

func (env Env) Clear() { os.Clearenv() }
func (env Env) Command(path string, args ...string) (cmd ICmd) {
	cmd = exec.Command(path, args...)
	return cmd
}
