package myerror

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
	ErrUnmarshal       = "unmarshal json"
	ErrMarshal         = "marshaling in json"
	ErrCheck           = "err check"
	ErrEncode          = "Encode"
	ErrInternal        = "err internal"
	UnknownReqId       = -1
)
