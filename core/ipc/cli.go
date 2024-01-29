package ipc

import (
	"log"

	"github.com/reshifr/play/core"
)

type CLI struct {
	handler core.HOS
}

func OpenCLI(handler core.HOS) (cli *CLI) {
	cli = &CLI{handler: handler}
	return cli
}

func (cli *CLI) Exec(bin string, args ...string) (output []byte, code int) {
	path := "/usr/bin/" + bin
	cmd := cli.handler.Command(path, args...)
	output, code = cli.execute(cmd)
	if code != core.CMD_EXIT_SUCCESS {
		log.Fatalf("Exec: can not execute '%v'.", cmd.String())
		return nil, code
	}
	return output, code
}

func (cli *CLI) execute(cmd core.IOSCmd) (output []byte, code int) {
	cli.handler.Clearenv()
	output, err := cmd.Output()
	if err != nil {
		if cmdErr, ok := err.(core.IOSCmdErr); ok {
			return nil, cmdErr.ExitCode()
		}
		return nil, core.CMD_EXIT_FAILURE
	}
	return output, core.CMD_EXIT_SUCCESS
}
