package application

import "geometricSolver/internals/util"

type LineAppInterface interface {
	ParallelismTwoLinesApp(body util.BodyHTTP) (util.BodyHTTP, error)
	PerpendicularTwoLinesApp(body util.BodyHTTP) (util.BodyHTTP, error)
	CornerTwoLinesApp(body util.BodyHTTP) (util.BodyHTTP, error)
	VerticalLineApp(body util.BodyHTTP) (util.BodyHTTP, error)
	HorizontalLineApp(body util.BodyHTTP) (util.BodyHTTP, error)
}

type LineApp struct {
}

func (l *LineApp) ParallelismTwoLinesApp(body util.BodyHTTP) (util.BodyHTTP, error) {
	return util.BodyHTTP{}, nil
}

func (l *LineApp) PerpendicularTwoLinesApp(body util.BodyHTTP) (util.BodyHTTP, error) {
	return util.BodyHTTP{}, nil
}

func (l *LineApp) CornerTwoLinesApp(body util.BodyHTTP) (util.BodyHTTP, error) {
	return util.BodyHTTP{}, nil
}

func (l *LineApp) VerticalLineApp(body util.BodyHTTP) (util.BodyHTTP, error) {
	return util.BodyHTTP{}, nil
}

func (l *LineApp) HorizontalLineApp(body util.BodyHTTP) (util.BodyHTTP, error) {
	return util.BodyHTTP{}, nil
}
