package myerror

type CheckErrorInterface interface {
	SetRequestIdUser(reqId int)
}

type CheckError struct {
	RequestId int
	Logger    MultiLoggerInterface
}

func (c *CheckError) SetRequestIdUser(reqId int) {
	c.RequestId = reqId
}
