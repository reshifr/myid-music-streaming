package cli

import (
	"log"
)

const (
	SUCCESS = 0
	FAILURE = 1
)

type LRU interface {
	Get(key string) (value string, ok bool)
	Add(key string, value string) (evicted bool)
}

type OSCmd interface {
	String() string
	Output() ([]byte, error)
}

type OSCodeError interface {
	ExitCode() int
}

type OSCmdHandler struct {
	Command  func(bin string, args ...string) OSCmd
	Clearenv func()
}

type Cmd struct {
	cache   LRU
	handler OSCmdHandler
}

func OpenCmd(
	cache LRU,
	handler OSCmdHandler,
) (cmd *Cmd) {
	cmd = &Cmd{
		cache:   cache,
		handler: handler,
	}
	return cmd
}

func (cmd *Cmd) Exec(
	bin string,
	args ...string,
) (output []byte, code int) {
	path, ok := cmd.locatePath(bin)
	if !ok {
		log.Fatalf("Exec: can not locate the '%v' path.", bin)
		return nil, FAILURE
	}
	osCmd := cmd.handler.Command(path, args...)
	output, code = cmd.executePath(osCmd)
	if code != SUCCESS {
		log.Fatalf("Exec: can not execute '%v'.", osCmd.String())
		return nil, code
	}
	return output, code
}

func (cmd *Cmd) locatePath(bin string) (path string, ok bool) {
	path, ok = cmd.cache.Get(bin)
	if ok {
		return path, true
	}
	cmd.handler.Clearenv()
	osCmd := cmd.handler.Command("/bin/which", bin)
	output, err := osCmd.Output()
	if err != nil {
		return "", false
	}
	length := len(output)
	if length == 0 {
		return "", false
	}
	path = string(output[:length-1])
	cmd.cache.Add(bin, path)
	return path, true
}

func (cmd *Cmd) executePath(osCmd OSCmd) (output []byte, code int) {
	cmd.handler.Clearenv()
	output, err := osCmd.Output()
	if err != nil {
		if codeErr, ok := err.(OSCodeError); ok {
			return nil, codeErr.ExitCode()
		}
		return nil, FAILURE
	}
	return output, SUCCESS
}
