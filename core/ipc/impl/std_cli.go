package ipc_impl

import "github.com/reshifr/myid-music-streaming/core/ipc"

type StdCLI[Env ipc.Env] struct {
	env Env
}

func NewStdCLI[Env ipc.Env](env Env) StdCLI[Env] {
	return StdCLI[Env]{env: env}
}

func (cli StdCLI[Env]) Exec(bin string, args ...string) ([]byte, int) {
	cmd := cli.env.Command(bin, args...)
	output, err := cmd.Output()
	if err != nil {
		if cmdErr, ok := err.(ipc.CmdError); ok {
			return output, cmdErr.ExitCode()
		}
	}
	return output, 0
}
