package api

import (
	"encoding/json"
	errPkg "geometricSolver/internals/myerror"
	"geometricSolver/internals/solver/application"
	"geometricSolver/internals/util"
	"github.com/valyala/fasthttp"
	"net/http"
)

type SolverApiInterface interface {
	GeomSolverHandler(ctx *fasthttp.RequestCtx)
}

type SolverApi struct {
	Application application.SolverAppInterface
	CheckErrors errPkg.CheckErrorInterface
	Logger      errPkg.MultiLoggerInterface
}

func (l *SolverApi) GeomSolverHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		l.Logger.Errorf("%s", errConvert.Error())
	}
	if reqId != errPkg.IntNil {
		l.CheckErrors.SetRequestIdUser(reqId)
	} else {
		l.CheckErrors.SetRequestIdUser(errPkg.UnknownReqId)
	}

	var reqBody util.BodyHTTP
	errUnmarshal := json.Unmarshal(ctx.Request.Body(), &reqBody)
	if errUnmarshal != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrUnmarshal))
		l.Logger.Errorf("%s, %s, requestId: %d", errPkg.ErrUnmarshal, errUnmarshal.Error(), reqId)
		return
	}

	responseBody, errIn := l.Application.GeomSolverApp(reqBody)
	errOut, resultOut, codeHTTP := l.CheckErrors.CheckErrorGeomSolver(errIn)
	if errOut != nil {
		switch errOut.Error() {
		case errPkg.ErrMarshal:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody([]byte(errPkg.ErrMarshal))
			return
		case errPkg.ErrCheck:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody(resultOut)
			return
		}
	}

	request, errResponse := json.Marshal(&responseBody)
	if errResponse != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		l.Logger.Errorf("%s, %s, requestId: %d", errPkg.ErrEncode, errResponse.Error(), reqId)
		return
	}

	ctx.Response.SetBody(request)
	json.NewEncoder(ctx)
	ctx.Response.SetStatusCode(http.StatusOK)

}
