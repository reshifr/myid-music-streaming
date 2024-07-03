package exec

type ExitError interface {
	ExitCode() int
	Error() string
}

type Cmd interface {
	Output() ([]byte, error)
}

type Exec interface {
	Command(name string, arg ...string) Cmd
}
