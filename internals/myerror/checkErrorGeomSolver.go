package myerror

import (
	"encoding/json"
	"net/http"
)

type CheckErrorInterface interface {
	SetRequestIdUser(reqId int)
	CheckErrorGeomSolver(inErr error) (error, []byte, int)
}

type CheckError struct {
	RequestId int
	Logger    MultiLoggerInterface
}

func (c *CheckError) SetRequestIdUser(reqId int) {
	c.RequestId = reqId
}

func (c *CheckError) CheckErrorGeomSolver(inErr error) (error, []byte, int) {
	if inErr != nil {
		errConvert, myError := ConvertErrorToMyErrors(inErr)
		if errConvert != nil {
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if errMarshal != nil {
				c.Logger.Errorf("%s, %v, requestId: %d", ErrMarshal, errMarshal, c.RequestId)
				return &MyErrors{
						ProjectTypeText: ErrMarshal,
						SourceText:      errMarshal.Error(),
					},
					nil, http.StatusInternalServerError
			}
			c.Logger.Errorf("%s, requestId: %d", inErr.Error(), c.RequestId)
			return &MyErrors{
					ProjectTypeText: ErrCheck,
					SourceText:      errConvert.Error(),
				},
				result, http.StatusInternalServerError

		}

		switch myError.ProjectTypeText {
		case NotFoundSolver:
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusConflict,
				Explain: ErrDB,
			})
			if errMarshal != nil {
				c.Logger.Errorf("%s, %v, requestId: %d", ErrMarshal, errMarshal, c.RequestId)
				return &MyErrors{
						ProjectTypeText: ErrMarshal,
						SourceText:      errMarshal.Error(),
					},
					nil, http.StatusInternalServerError
			}
			c.Logger.Warnf("Type: %s, source: %s, requestId: %d", myError.ProjectTypeText, myError.SourceText, c.RequestId)
			return &MyErrors{
					ProjectTypeText: ErrCheck,
					SourceText:      myError.ProjectTypeText,
				},
				result, http.StatusConflict
		default:
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if errMarshal != nil {
				c.Logger.Errorf("%s, %v, requestId: %d", ErrMarshal, errMarshal, c.RequestId)
				return &MyErrors{
						ProjectTypeText: ErrMarshal,
						SourceText:      errMarshal.Error(),
					},
					nil, http.StatusInternalServerError
			}
			c.Logger.Errorf("%s, requestId: %d", inErr.Error(), c.RequestId)
			return &MyErrors{
					ProjectTypeText: ErrCheck,
					SourceText:      myError.ProjectTypeText,
				},
				result, http.StatusInternalServerError

		}

	}
	return nil, nil, IntNil
}
