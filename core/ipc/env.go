package ipc

type Env interface {
	Command(bin string, args ...string) (cmd Cmd)
}

type Cmd interface {
	String() (cmd string)
	Output() (output []byte, err error)
}

type CmdError interface {
	ExitCode() (code int)
}
