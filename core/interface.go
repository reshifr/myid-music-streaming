package core

type OSCmd interface {
	String() (cmd string)
	Output() (output []byte, err error)
}

type OSCmdError interface {
	ExitCode() (code int)
}

type LRU[K comparable, V any] interface {
	Get(key K) (value V, ok bool)
	Insert(key K, value V)
	Contains(key K) (ok bool)
}
