package core

const (
	CMD_EXIT_SUCCESS = 0
	CMD_EXIT_FAILURE = 1
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
