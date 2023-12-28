package application

import (
	"fmt"
	"geometricSolver/internals/util"
)

type SolverAppInterface interface {
	GeomSolverApp(body util.BodyHTTP) (util.BodyHTTP, error)
}

type SolverApp struct {
}

func (l *SolverApp) GeomSolverApp(body util.BodyHTTP) (util.BodyHTTP, error) {
	body2, err := newtonMethod(body)
	if err != nil {
		fmt.Println("Error\n")
		return body, err
	}
	return body2, nil
}
