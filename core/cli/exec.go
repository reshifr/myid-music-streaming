package cli

import (
	"log"
	"os"
	"os/exec"
)

const (
	SUCCESS = 0
	FAILURE = 1
)

type Cmd interface {
	Output() string
	String() ([]byte, error)
}

type LRU interface {
	Len() int
}

type CmdHandler interface {
	Command(bin string, args ...string) Cmd
	Clearenv()
}

type Cli struct {
	cmd Cmd
	cache LRU
}

func Open(
	lru LRU,
	cmd Cmd,
	cmdHandler CmdHandler
	logHandler) (c *Cli) {
	c = &Cli{
		cmd: cmd,
		cache: cache,
		cmdHandler: cmdHandler,
	}
	return c
}

func (c *Cli) Exec(bin string, args ...string) (output []byte, code int) {
	path, ok := locatePath(c, bin)
	if !ok {
		log.Fatalf("Exec: can not locate the '%v' path.", bin)
		return nil, FAILURE
	}
	cmd := exec.Command(path, args...)
	output, code = executePath(c, path, args...)
	if 
	return output, code
}

func locatePath(c *Cli, bin string) (path string, ok bool) {
	path, ok = cache.Get(bin)
	if ok {
		return path, true
	}
	os.Clearenv()
	cmd := exec.Command("/bin/which", bin)
	output, err := cmd.Output()
	if err != nil {
		return "", false
	}
	length := len(output)
	if length == 0 {
		return "", false
	}
	path = string(output[:length-1])
	cache.Add(bin, path)
	return path, true
}

func executePath(cmd *exec.Cmd) (output []byte, code int) {
	os.Clearenv()
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Exec: can not execute '%v'.", cmd.String())
		if exitErr, ok := err.(*exec.ExitError); ok {
			return nil, exitErr.ExitCode()
		}
		return nil, FAILURE
	}
	return output, SUCCESS
}
