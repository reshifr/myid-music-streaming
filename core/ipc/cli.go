package ipc

import (
	"github.com/reshifr/play/core"
)

type CLI[E core.IEnv] struct {
	env E
}

func OpenCLI[E core.IEnv](env E) CLI[E] {
	return CLI[E]{env: env}
}

func (cli *CLI[E]) Exec(bin string, args ...string) ([]byte, *core.Error) {
	path := "/usr/bin/" + bin
	cmd := cli.env.Command(path, args...)
	output, code := cli.execute(cmd)
	var cerr *core.Error = nil
	if code != core.CmdExitOk {
		cerr = core.NewErrorf(
			code,
			"ipc.CLI.Exec(): error execute '%v'.",
			cmd.String(),
		)
	}
	return output, cerr
}

func (cli *CLI[E]) execute(cmd core.ICmd) ([]byte, int) {
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
