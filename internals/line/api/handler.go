package api

import (
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

}
