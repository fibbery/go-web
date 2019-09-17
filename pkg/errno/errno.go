package errno

import "fmt"

type Errno struct {
	Code    int
	Message string
}

func (e Errno) Error() string {
	return e.Message
}

type Err struct {
	Code    int
	Message string
	Err     error
}

func (e Err) Error() string {
	return fmt.Sprintf("Err - Code : %d, Message : %s, Error: %s", e.Code, e.Message, e.Err)
}

func (e *Err) add(message string) *Err {
	e.Message += " " + message
	return e
}

func (e *Err) addf(format string, args ...interface{}) *Err {
	e.Message += " " + fmt.Sprintf(format, args...)
	return e
}

func New(errno *Errno, err error) *Err {
	return &Err{
		Code:    errno.Code,
		Message: errno.Message,
		Err:     err,
	}
}

func DecodeError(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}
	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Err.Error()
	case *Errno:
		return typed.Code, typed.Message
	default:
	}
	return InternalServerError.Code, err.Error()
}
