package ipc

type ICLI interface {
	Exec(bin string, args ...string) (output []byte, code int)
}
