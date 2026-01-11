package ecode

import (
	"fmt"

	"google.golang.org/grpc/status"
)

type Codes interface {
	Error() string
	Code() int32
	Message() string
}

type eCode struct {
	code    int32
	message string
}

func (e *eCode) Error() string {
	if e == nil || e.code == 0 {
		return ""
	}
	return fmt.Sprintf("code=%d,msg=%s", e.code, e.message)
}

func (e *eCode) Code() int32 {
	if e == nil {
		return 0
	}
	return e.code
}

func (e *eCode) Message() string {
	if e == nil {
		return ""
	}
	return e.message
}

func Error(err error) Codes {
	if err == nil {
		return nil
	}

	if e, ok := err.(Codes); ok {
		return e
	}

	if e, ok := status.FromError(err); ok {
		return &eCode{int32(e.Code()), e.Message()}
	}

	return &eCode{OriginError, err.Error()}
}

func Errorf(code int32, opts ...interface{}) Codes {
	if code == 0 {
		return nil
	}
	e := &eCode{
		code:    code,
		message: "",
	}
	if len(opts) == 1 {
		e.message = opts[0].(string)
	} else if len(opts) > 1 {
		e.message = fmt.Sprintf(opts[0].(string), opts[1:]...)
	}
	return e
}
