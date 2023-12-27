package api

import (
	"encoding/json"
	"geometricSolver/internals/line/application"
	errPkg "geometricSolver/internals/myerror"
	"geometricSolver/internals/util"
	"github.com/valyala/fasthttp"
	"net/http"
)

type LineApiInterface interface {
	ParallelismTwoLinesHandler(ctx *fasthttp.RequestCtx)
	PerpendicularTwoLinesHandler(ctx *fasthttp.RequestCtx)
	CornerTwoLinesHandler(ctx *fasthttp.RequestCtx)
	VerticalLineHandler(ctx *fasthttp.RequestCtx)
	HorizontalLineHandler(ctx *fasthttp.RequestCtx)
}

type LineApi struct {
	Application application.LineAppInterface
	CheckErrors errPkg.CheckErrorInterface
	Logger      errPkg.MultiLoggerInterface
}

func (l *LineApi) ParallelismTwoLinesHandler(ctx *fasthttp.RequestCtx) {
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

	responseBody, errIn := l.Application.ParallelismTwoLinesApp(reqBody)
	errOut, resultOut, codeHTTP := l.CheckErrors.CheckErrorParallelismTwoLines(errIn)
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

	//responseBody.Points = append(responseBody.Points, util.Point{X: 5.6, Y: 6.9})
	//responseBody.Points = append(responseBody.Points, util.Point{X: 9.6, Y: 6.9})
	//responseBody.Points = append(responseBody.Points, util.Point{X: 5.6, Y: 13.9})
	//responseBody.Points = append(responseBody.Points, util.Point{X: 9.6, Y: 18.9})
	//
	//responseBody.Lines = append(responseBody.Lines, util.PairNumber{First: 0, Second: 1})
	//responseBody.Lines = append(responseBody.Lines, util.PairNumber{First: 2, Second: 3})
	//
	//responseBody.HorizontLine = append(responseBody.HorizontLine, 0)

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

func (l *LineApi) PerpendicularTwoLinesHandler(ctx *fasthttp.RequestCtx) {
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
	ctx.Response.SetStatusCode(http.StatusNotImplemented) //TODO(N): remake
}

func (l *LineApi) CornerTwoLinesHandler(ctx *fasthttp.RequestCtx) {
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
	ctx.Response.SetStatusCode(http.StatusNotImplemented) //TODO(N): remake
}

func (l *LineApi) VerticalLineHandler(ctx *fasthttp.RequestCtx) {
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
	ctx.Response.SetStatusCode(http.StatusNotImplemented) //TODO(N): remake
}

func (l *LineApi) HorizontalLineHandler(ctx *fasthttp.RequestCtx) {
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
	ctx.Response.SetStatusCode(http.StatusNotImplemented) //TODO(N): remake
}
