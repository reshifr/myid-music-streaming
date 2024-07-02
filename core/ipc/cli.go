package ipc

type CLI interface {
	Exec(bin string, args ...string) (output []byte, code int)
}
