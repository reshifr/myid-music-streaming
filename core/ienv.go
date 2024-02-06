package core

const (
	CmdExitOk = iota
	CmdExitFail
)

type IEnv interface {
	Clear()
	Command(path string, args ...string) (cmd ICmd)
}

type ICmd interface {
	String() (cmd string)
	Output() (output []byte, err error)
}

type ICmdError interface {
	ExitCode() (code int)
}
