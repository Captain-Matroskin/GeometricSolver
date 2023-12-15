package myerror

import (
	"errors"
)

type MultiLoggerInterface interface {
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Sync() error
}

type MyErrors struct {
	SourceText      string
	ProjectTypeText string
	Way             string
}

func ConvertErrorToMyErrors(err error) (error, *MyErrors) {
	var myErrors *MyErrors
	switch {
	case errors.As(err, &myErrors):
		return nil, myErrors
	default:
		return &MyErrors{
			ProjectTypeText: ErrNotMyErrors,
		}, nil
	}
}

func (e *MyErrors) Error() string {
	return e.ProjectTypeText
}

type ResultError struct {
	Status  int    `json:"status"`
	Explain string `json:"explain,omitempty"`
}

// Error of server
const (
	ErrDB              = "database is not responding"
	ErrAtoi            = "func Atoi convert string in int"
	IntNil             = 0
	ErrNotStringAndInt = "expected type string or int"
	ErrNotMyErrors     = "expected type MyErrors"
	ErrUnmarshal       = "unmarshal json"
	ErrMarshal         = "marshaling in json"
	ErrCheck           = "err check"
	ErrEncode          = "Encode"
	ErrInternal        = "err internal"
	UnknownReqId       = -1
	UnknowTypeError    = "an unexpected type of error. Not equal MyErrors"
)
