package core

import (
	"fmt"
)

type Error struct {
	Msg  string
	Code int
}

func ThrowError(code int, msg string) *Error {
	return &Error{
		Msg:  msg,
		Code: code,
	}
}

func ThrowErrorf(code int, format string, args ...any) *Error {
	return &Error{
		Msg:  fmt.Sprintf(format, args...),
		Code: code,
	}
}
