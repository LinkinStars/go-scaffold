package mistake

import (
	"bytes"
	"fmt"
	"strings"
)

// Error is error with info
type Error struct {
	Code    int
	Reason  string
	Message string
	Err     error
	Stack   string
}

// New create error
func New(code int, reason, message string) *Error {
	return &Error{Code: code, Reason: reason, Message: message}
}

// Error return error with info
func (e *Error) Error() string {
	return e.Message
}

// WithError with original error
func (e *Error) WithError(err error) *Error {
	e.Err = err
	return e
}

// WithStack with stack
func (e *Error) WithStack() *Error {
	e.Stack = LogStack(2, 0)
	return e
}

func (e *Error) Format(state fmt.State, verb rune) {
	switch verb {
	case 'v':
		str := bytes.NewBuffer([]byte{})
		str.WriteString(fmt.Sprintf("code: %d, ", e.Code))
		str.WriteString("reason: ")
		str.WriteString(e.Reason + ", ")
		str.WriteString("message: ")
		str.WriteString(e.Message)
		if e.Err != nil {
			str.WriteString(", error: ")
			str.WriteString(e.Err.Error())
		}
		if len(e.Stack) > 0 {
			str.WriteString("\n")
			str.WriteString(e.Stack)
		}
		fmt.Fprintf(state, "%s", strings.Trim(str.String(), "\r\n\t"))
	default:
		fmt.Fprintf(state, e.Message)
	}
}
