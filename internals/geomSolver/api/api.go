package api

import (
	"geometricSolver/internals/geomSolver/application"
	errPkg "geometricSolver/internals/myerror"
	"github.com/valyala/fasthttp"
)

type GeomSolverApiInterface interface {
	EqualTwoPointsHandler(ctx *fasthttp.RequestCtx)
	DistanceBetweenTwoPointsHandler(ctx *fasthttp.RequestCtx)
	FixationPointHandler(ctx *fasthttp.RequestCtx)
	ParallelismTwoLinesHandler(ctx *fasthttp.RequestCtx)
	PerpendicularTwoLinesHandler(ctx *fasthttp.RequestCtx)
	CornerTwoLinesHandler(ctx *fasthttp.RequestCtx)
	VerticalityLineHandler(ctx *fasthttp.RequestCtx)
	HorizontalLineHandler(ctx *fasthttp.RequestCtx)
	BelongingPointOfLineHandler(ctx *fasthttp.RequestCtx)
}

// уровень апи
type GeomSolverApi struct {
	Application application.GeomSolverAppInterface
	CheckErrors errPkg.CheckErrorInterface
	Logger      errPkg.MultiLoggerInterface
}

func (g *GeomSolverApi) EqualTwoPointsHandler(ctx *fasthttp.RequestCtx) {

}

func (g *GeomSolverApi) DistanceBetweenTwoPointsHandler(ctx *fasthttp.RequestCtx) {

}

func (g *GeomSolverApi) FixationPointHandler(ctx *fasthttp.RequestCtx) {

}

func (g *GeomSolverApi) ParallelismTwoLinesHandler(ctx *fasthttp.RequestCtx) {

}

func (g *GeomSolverApi) PerpendicularTwoLinesHandler(ctx *fasthttp.RequestCtx) {

}

func (g *GeomSolverApi) CornerTwoLinesHandler(ctx *fasthttp.RequestCtx) {

}

func (g *GeomSolverApi) VerticalityLineHandler(ctx *fasthttp.RequestCtx) {

}

func (g *GeomSolverApi) HorizontalLineHandler(ctx *fasthttp.RequestCtx) {

}

func (g *GeomSolverApi) BelongingPointOfLineHandler(ctx *fasthttp.RequestCtx) {

}
