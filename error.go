package errorx

import (
	"fmt"
)

type Error interface {
	Error() string  // error interface
	String() string // Stringer interface

	Code() int
	Message() string
}

type ErrorPayload struct {
	CodePayload    int    `json:"code"`
	MessagePayload string `json:"message"`
}

func New(code int, message string) Error {
	return ErrorPayload{
		CodePayload:    code,
		MessagePayload: message,
	}
}

func NewFromErr(code int, err error) Error {
	return ErrorPayload{
		CodePayload:    code,
		MessagePayload: err.Error(),
	}
}

func NewErrorPayloadFromError(errx Error) *ErrorPayload {
	return &ErrorPayload{
		CodePayload:    errx.Code(),
		MessagePayload: errx.Message(),
	}
}

func (thisRef ErrorPayload) Code() int {
	return thisRef.CodePayload
}

func (thisRef ErrorPayload) Message() string {
	return thisRef.MessagePayload
}

func (thisRef ErrorPayload) String() string {
	return fmt.Sprintf("message: %s, code: %d", thisRef.MessagePayload, thisRef.CodePayload)
}

func (thisRef ErrorPayload) Error() string {
	return thisRef.String()
}
