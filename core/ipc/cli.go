package ipc

import (
	"log"

	"github.com/reshifr/play/core"
)

type CLI struct {
	env core.IEnv
}

func OpenCLI(env core.IEnv) (cli *CLI) {
	cli = &CLI{env: env}
	return cli
}

func (cli *CLI) Exec(bin string, args ...string) (output []byte, code int) {
	path := "/usr/bin/" + bin
	cmd := cli.env.Command(path, args...)
	output, code = cli.execute(cmd)
	if code != core.CMD_EXIT_SUCCESS {
		log.Fatalf("Exec: can not execute '%v'.", cmd.String())
	}
	return output, code
}

func (cli *CLI) execute(cmd core.ICmd) (output []byte, code int) {
	cli.env.Clear()
	output, err := cmd.Output()
	code = core.CMD_EXIT_SUCCESS
	if err != nil {
		code = core.CMD_EXIT_FAILURE
		if cmdErr, ok := err.(core.ICmdErr); ok {
			code = cmdErr.ExitCode()
		}
	}
	return output, code
}
