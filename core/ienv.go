package core

const (
	CmdExitOk   = 0
	CmdExitFail = 1
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
