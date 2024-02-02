package ipc

import (
	"github.com/reshifr/play/core"
)

type CLI[Env core.IEnv] struct {
	env Env
}

func OpenCLI[Env core.IEnv](env Env) *CLI[Env] {
	return &CLI[Env]{env: env}
}

func (cli *CLI[Env]) Exec(bin string, args ...string) ([]byte, *core.Error) {
	path := "/usr/bin/" + bin
	cmd := cli.env.Command(path, args...)
	output, code := cli.execute(cmd)
	var coreErr *core.Error = nil
	if code != core.CmdExitOk {
		coreErr = core.ThrowErrorf(
			code,
			"ipc.CLI.Exec(): error execute '%v'.",
			cmd.String(),
		)
	}
	return output, coreErr
}

func (cli *CLI[Env]) execute(cmd core.ICmd) ([]byte, int) {
	cli.env.Clear()
	output, err := cmd.Output()
	code := core.CmdExitOk
	if err != nil {
		code = core.CmdExitFail
		if cmdErr, ok := err.(core.ICmdError); ok {
			code = cmdErr.ExitCode()
		}
	}
	return output, code
}
