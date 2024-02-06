package core

import (
	"fmt"
)

type Error struct {
	Msg  string
	Err  *Error
	Code int
}

func NewError(code int, msg string) *Error {
	return &Error{
		Msg:  msg,
		Err:  nil,
		Code: code,
	}
}

func NewErrorf(code int, format string, args ...any) *Error {
	return &Error{
		Msg:  fmt.Sprintf(format, args...),
		Err:  nil,
		Code: code,
	}
}

func NewChainError(err *Error, code int, msg string) *Error {
	return &Error{
		Msg:  msg,
		Err:  err,
		Code: code,
	}
}

func NewChainErrorf(err *Error, code int, format string, args ...any) *Error {
	return &Error{
		Msg:  fmt.Sprintf(format, args...),
		Err:  err,
		Code: code,
	}
}
