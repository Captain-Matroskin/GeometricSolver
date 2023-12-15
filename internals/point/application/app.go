package application

import "geometricSolver/internals/util"

type PointAppInterface interface {
	EqualTwoPointsApp(body util.BodyHTTP) (util.BodyHTTP, error)
	DistanceBetweenTwoPointsApp(body util.BodyHTTP) (util.BodyHTTP, error)
	FixationPointApp(body util.BodyHTTP) (util.BodyHTTP, error)
	BelongingPointOfLineApp(body util.BodyHTTP) (util.BodyHTTP, error)
}

type PointApp struct {
}

func (p *PointApp) EqualTwoPointsApp(body util.BodyHTTP) (util.BodyHTTP, error) {
	return util.BodyHTTP{}, nil
}

func (p *PointApp) DistanceBetweenTwoPointsApp(body util.BodyHTTP) (util.BodyHTTP, error) {
	return util.BodyHTTP{}, nil
}

func (p *PointApp) FixationPointApp(body util.BodyHTTP) (util.BodyHTTP, error) {
	return util.BodyHTTP{}, nil
}

func (p *PointApp) BelongingPointOfLineApp(body util.BodyHTTP) (util.BodyHTTP, error) {
	return util.BodyHTTP{}, nil
}
