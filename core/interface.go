package core

type ICmd interface {
	String() (cmd string)
	Output() (output []byte, err error)
}

type ICmdErr interface {
	ExitCode() (code int)
}

type IEnv interface {
	Clear()
	Command(path string, args ...string) (cmd ICmd)
}
