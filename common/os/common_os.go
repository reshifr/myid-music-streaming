package os

import "os"

type CommonOS struct{}

func (CommonOS) Open(name string) (File, error) { return os.Open(name) }
