package ipc

const (
	ExitSuccess = iota
	ExitFailure
)

type CLI interface {
	Exec(bin string, args ...string) (output []byte, code int)
}
