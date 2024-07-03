package exec

import "os/exec"

type CommonExec struct{}

func (CommonExec) Command(name string, arg ...string) Cmd { return exec.Command(name, arg...) }
