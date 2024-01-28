package ipc

import (
	"log"

	"github.com/reshifr/play/core"
)

const (
	CLI_EXIT_SUCCESS = 0
	CLI_EXIT_FAILURE = 1
)

type CLI struct {
	handler core.OSHandler
}

func OpenCLI(handler core.OSHandler) (cli *CLI) {
	cli = &CLI{handler: handler}
	return cli
}

func (cli *CLI) Exec(bin string, args ...string) (output []byte, code int) {
	path := "/usr/bin/" + bin
	cmd := cli.handler.Command(path, args...)
	output, code = cli.executePath(cmd)
	if code != CLI_EXIT_SUCCESS {
		log.Fatalf("Exec: can not execute '%v'.", cmd.String())
		return nil, code
	}
	return output, code
}

func (cli *CLI) executePath(cmd core.OSCmd) (output []byte, code int) {
	cli.handler.Clearenv()
	output, err := cmd.Output()
	if err != nil {
		if cmdErr, ok := err.(core.OSCmdError); ok {
			return nil, cmdErr.ExitCode()
		}
		return nil, CLI_EXIT_FAILURE
	}
	return output, CLI_EXIT_SUCCESS
}
