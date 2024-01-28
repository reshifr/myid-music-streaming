package core

type IOSCmd interface {
	String() (cmd string)
	Output() (output []byte, err error)
}

type IOSCmdErr interface {
	ExitCode() (code int)
}

type HOS struct {
	Command  func(path string, args ...string) (cmd IOSCmd)
	Clearenv func()
}
