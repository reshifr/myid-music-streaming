package ipc

import "github.com/reshifr/myid-music-streaming/common/exec"

type CoreCLI[Exec exec.Exec] struct {
	exec Exec
}

func NewCoreCLI[Exec exec.Exec](exec Exec) CoreCLI[Exec] {
	return CoreCLI[Exec]{exec}
}

func (c CoreCLI[Exec]) Exec(bin string, args ...string) ([]byte, int) {
	cmd := c.exec.Command("/usr/bin/env", append([]string{bin}, args...)...)
	o, err := cmd.Output()
	if err != nil {
		if cmdErr, ok := err.(exec.ExitError); ok {
			return o, cmdErr.ExitCode()
		}
	}
	return o, ExitSuccess
}
