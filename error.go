package errorx

import (
	"fmt"
)

type Error interface {
	Error() string  // error interface
	String() string // Stringer interface

	Code() int
	Message() string
	Data() interface{}
}

type ErrorPayload struct {
	CodePayload    int         `json:"code"`
	MessagePayload string      `json:"message,omitempty"`
	DataPlayload   interface{} `json:"data,omitempty"`
}

func New(code int, message string) Error {
	return ErrorPayload{
		CodePayload:    code,
		MessagePayload: message,
	}
}

func NewWithData(code int, message string, data interface{}) Error {
	return ErrorPayload{
		CodePayload:    code,
		MessagePayload: message,
		DataPlayload:   data,
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
		DataPlayload:   errx.Data(),
	}
}

func (thisRef ErrorPayload) Code() int {
	return thisRef.CodePayload
}

func (thisRef ErrorPayload) Message() string {
	return thisRef.MessagePayload
}

func (thisRef ErrorPayload) Data() interface{} {
	return thisRef.DataPlayload
}

func (thisRef ErrorPayload) String() string {
	if thisRef.DataPlayload != nil {
		return fmt.Sprintf("code: %d, message: %s, data: %v", thisRef.CodePayload, thisRef.MessagePayload, thisRef.DataPlayload)
	}

	return fmt.Sprintf("code: %d, message: %s", thisRef.CodePayload, thisRef.MessagePayload)
}

func (thisRef ErrorPayload) Error() string {
	return thisRef.String()
}
