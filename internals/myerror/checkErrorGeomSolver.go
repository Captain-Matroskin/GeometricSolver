package myerror

type CheckErrorInterface interface {
}

type CheckError struct {
	RequestId int
	Logger    MultiLoggerInterface
}
