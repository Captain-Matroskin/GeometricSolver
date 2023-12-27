package application

import (
	"fmt"
	"geometricSolver/internals/util"
)

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
	body2, err := newtonMethod(body)
	if err != nil {
		fmt.Println("Error\n")
		return body, err
	}
	return body2, nil
}

func (l *LineApp) PerpendicularTwoLinesApp(body util.BodyHTTP) (util.BodyHTTP, error) {
	body, err := newtonMethod(body)
	if err != nil {
		fmt.Println("Error\n")
	}
	return body, nil
}

func (l *LineApp) CornerTwoLinesApp(body util.BodyHTTP) (util.BodyHTTP, error) {
	body, err := newtonMethod(body)
	if err != nil {
		fmt.Println("Error\n")
	}
	return body, nil
}

func (l *LineApp) VerticalLineApp(body util.BodyHTTP) (util.BodyHTTP, error) {
	body, err := newtonMethod(body)
	if err != nil {
		fmt.Println("Error\n")
	}
	return body, nil
}

func (l *LineApp) HorizontalLineApp(body util.BodyHTTP) (util.BodyHTTP, error) {
	body, err := newtonMethod(body)
	if err != nil {
		fmt.Println("Error\n")
	}
	return body, nil
}
