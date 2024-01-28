package core

type OSHandler struct {
	Command  func(path string, args ...string) (cmd OSCmd)
	Clearenv func()
}
