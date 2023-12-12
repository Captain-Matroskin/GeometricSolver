package api

import (
	errPkg "geometricSolver/internals/myerror"
	"geometricSolver/internals/point/application"
	"geometricSolver/internals/util"
	"github.com/valyala/fasthttp"
	"net/http"
)

type PointApiInterface interface {
	EqualTwoPointsHandler(ctx *fasthttp.RequestCtx)
	DistanceBetweenTwoPointsHandler(ctx *fasthttp.RequestCtx)
	FixationPointHandler(ctx *fasthttp.RequestCtx)
	BelongingPointOfLineHandler(ctx *fasthttp.RequestCtx)
}

type PointApi struct {
	Application application.PointAppInterface
	CheckErrors errPkg.CheckErrorInterface
	Logger      errPkg.MultiLoggerInterface
}

func (p *PointApi) EqualTwoPointsHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		p.Logger.Errorf("%s", errConvert.Error())
	}
	if reqId != errPkg.IntNil {
		p.CheckErrors.SetRequestIdUser(reqId)
	} else {
		p.CheckErrors.SetRequestIdUser(errPkg.UnknownReqId)
	}

}

func (p *PointApi) DistanceBetweenTwoPointsHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		p.Logger.Errorf("%s", errConvert.Error())
	}
	if reqId != errPkg.IntNil {
		p.CheckErrors.SetRequestIdUser(reqId)
	} else {
		p.CheckErrors.SetRequestIdUser(errPkg.UnknownReqId)
	}

}

func (p *PointApi) FixationPointHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		p.Logger.Errorf("%s", errConvert.Error())
	}
	if reqId != errPkg.IntNil {
		p.CheckErrors.SetRequestIdUser(reqId)
	} else {
		p.CheckErrors.SetRequestIdUser(errPkg.UnknownReqId)
	}

}

func (p *PointApi) BelongingPointOfLineHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		p.Logger.Errorf("%s", errConvert.Error())
	}
	if reqId != errPkg.IntNil {
		p.CheckErrors.SetRequestIdUser(reqId)
	} else {
		p.CheckErrors.SetRequestIdUser(errPkg.UnknownReqId)
	}

}
