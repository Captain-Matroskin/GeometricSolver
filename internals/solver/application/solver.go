package application

import (
	"fmt"
	errPkg "geometricSolver/internals/myerror"
	"geometricSolver/internals/util"
	"gonum.org/v1/gonum/diff/fd"
	"gonum.org/v1/gonum/mat"
	"math"
)

func b_creater(body util.BodyHTTP, x []float64) *mat.VecDense {
	nEqual := len(body.EqualTwoPoints)
	nFix := len(body.FixationPoint)
	nDist := len(body.DistanceBetweenPoints)
	nBelong := len(body.BelongOfLine)
	nParall := len(body.ParallelTwoLines)
	nPerpen := len(body.PerpenTwoLines)
	nCorner := len(body.CornerTwoLines)
	nVert := len(body.VerticalLine)
	nHoriz := len(body.HorizontLine)

	size_constraint := 2*nEqual + 2*nFix + nDist + nBelong + nParall + nPerpen + nCorner + nVert + nHoriz
	size_matrix := 2*len(body.Points) + size_constraint

	b := mat.NewVecDense(size_matrix, make([]float64, size_matrix))

	for i := 0; i < size_matrix; i++ {
		b.SetVec(i, 0.0)
	}

	for i, value := range body.EqualTwoPoints {
		first := value.First
		second := value.Second
		x1 := body.Points[first].X
		y1 := body.Points[first].Y
		x2 := body.Points[second].X
		y2 := body.Points[second].Y
		dl_x1 := x[size_constraint+2*first]
		dl_y1 := x[size_constraint+2*first+1]
		dl_x2 := x[size_constraint+2*second]
		dl_y2 := x[size_constraint+2*second+1]

		b.SetVec(2*i, x2+dl_x2-x1-dl_x1)
		b.SetVec(2*i+1, y2+dl_y2-y1-dl_y1)
	}

	for i, value := range body.FixationPoint {
		num_point := value
		x1 := body.Points[num_point].X
		y1 := body.Points[num_point].Y
		xc := x1
		yc := y1
		dl_x1 := x[size_constraint+2*num_point]
		dl_y1 := x[size_constraint+2*num_point+1]

		b.SetVec(2*nEqual+2*i, x1+dl_x1-xc)
		b.SetVec(2*nEqual+2*i+1, y1+dl_y1-yc)
	}

	for i, value := range body.DistanceBetweenPoints {
		first := value.First
		second := value.Second
		d := value.Value
		x1 := body.Points[first].X
		y1 := body.Points[first].Y
		x2 := body.Points[second].X
		y2 := body.Points[second].Y
		dl_x1 := x[size_constraint+2*first]
		dl_y1 := x[size_constraint+2*first+1]
		dl_x2 := x[size_constraint+2*second]
		dl_y2 := x[size_constraint+2*second+1]

		b.SetVec(2*nEqual+2*nFix+i, (x2+dl_x2-x1-dl_x1)*(x2+dl_x2-x1-dl_x1)+(y2+dl_y2-y1-dl_y1)*(y2+dl_y2-y1-dl_y1)-d*d)
	}

	for i, value := range body.BelongOfLine {
		num_point := value.First
		num_line := value.Second
		xp := body.Points[num_point].X
		yp := body.Points[num_point].Y
		x1 := body.Points[body.Lines[num_line].First].X
		y1 := body.Points[body.Lines[num_line].First].Y
		x2 := body.Points[body.Lines[num_line].Second].X
		y2 := body.Points[body.Lines[num_line].Second].Y
		dl_xp := x[size_constraint+2*num_point]
		dl_yp := x[size_constraint+2*num_point+1]
		dl_x1 := x[size_constraint+2*body.Lines[num_line].First]
		dl_y1 := x[size_constraint+2*body.Lines[num_line].First+1]
		dl_x2 := x[size_constraint+2*body.Lines[num_line].Second]
		dl_y2 := x[size_constraint+2*body.Lines[num_line].Second+1]

		b.SetVec(2*nEqual+2*nFix+nDist+i, (xp+dl_xp-x1-dl_x1)*(y2+dl_y2-yp-dl_yp)-(x2+dl_x2-xp-dl_xp)*(yp+dl_yp-y1-dl_y1))
	}

	for i, value := range body.ParallelTwoLines {
		num_line1 := value.First
		num_line2 := value.Second
		line1_first_point := body.Lines[num_line1].First
		line1_second_point := body.Lines[num_line1].Second
		line2_first_point := body.Lines[num_line2].First
		line2_second_point := body.Lines[num_line2].Second
		x1 := body.Points[line1_first_point].X
		y1 := body.Points[line1_first_point].Y
		x2 := body.Points[line1_second_point].X
		y2 := body.Points[line1_second_point].Y
		x3 := body.Points[line2_first_point].X
		y3 := body.Points[line2_first_point].Y
		x4 := body.Points[line2_second_point].X
		y4 := body.Points[line2_second_point].Y
		dl_x1 := x[size_constraint+2*body.Lines[num_line1].First]
		dl_y1 := x[size_constraint+2*body.Lines[num_line1].First+1]
		dl_x2 := x[size_constraint+2*body.Lines[num_line1].Second]
		dl_y2 := x[size_constraint+2*body.Lines[num_line1].Second+1]
		dl_x3 := x[size_constraint+2*body.Lines[num_line2].First]
		dl_y3 := x[size_constraint+2*body.Lines[num_line2].First+1]
		dl_x4 := x[size_constraint+2*body.Lines[num_line2].Second]
		dl_y4 := x[size_constraint+2*body.Lines[num_line2].Second+1]

		b.SetVec(2*nEqual+2*nFix+nDist+nBelong+i, (x2+dl_x2-x1-dl_x1)*(y4+dl_y4-y3-dl_y3)-(x4+dl_x4-x3-dl_x3)*(y2+dl_y2-y1-dl_y1))
	}

	for i, value := range body.PerpenTwoLines {
		num_line1 := value.First
		num_line2 := value.Second
		line1_first_point := body.Lines[num_line1].First
		line1_second_point := body.Lines[num_line1].Second
		line2_first_point := body.Lines[num_line2].First
		line2_second_point := body.Lines[num_line2].Second
		x1 := body.Points[line1_first_point].X
		y1 := body.Points[line1_first_point].Y
		x2 := body.Points[line1_second_point].X
		y2 := body.Points[line1_second_point].Y
		x3 := body.Points[line2_first_point].X
		y3 := body.Points[line2_first_point].Y
		x4 := body.Points[line2_second_point].X
		y4 := body.Points[line2_second_point].Y
		dl_x1 := x[size_constraint+2*body.Lines[num_line1].First]
		dl_y1 := x[size_constraint+2*body.Lines[num_line1].First+1]
		dl_x2 := x[size_constraint+2*body.Lines[num_line1].Second]
		dl_y2 := x[size_constraint+2*body.Lines[num_line1].Second+1]
		dl_x3 := x[size_constraint+2*body.Lines[num_line2].First]
		dl_y3 := x[size_constraint+2*body.Lines[num_line2].First+1]
		dl_x4 := x[size_constraint+2*body.Lines[num_line2].Second]
		dl_y4 := x[size_constraint+2*body.Lines[num_line2].Second+1]

		b.SetVec(2*nEqual+2*nFix+nDist+nBelong+nParall+i, (x2+dl_x2-x1-dl_x1)*(x4+dl_x4-x3-dl_x3)+(y2+dl_y2-y1-dl_y1)*(y4+dl_y4-y3-dl_y3))
	}

	for i, value := range body.CornerTwoLines {
		num_line1 := value.First
		num_line2 := value.Second
		//corner := value.
		line1_first_point := body.Lines[num_line1].First
		line1_second_point := body.Lines[num_line1].Second
		line2_first_point := body.Lines[num_line2].First
		line2_second_point := body.Lines[num_line2].Second
		x1 := body.Points[line1_first_point].X
		y1 := body.Points[line1_first_point].Y
		x2 := body.Points[line1_second_point].X
		y2 := body.Points[line1_second_point].Y
		x3 := body.Points[line2_first_point].X
		y3 := body.Points[line2_first_point].Y
		x4 := body.Points[line2_second_point].X
		y4 := body.Points[line2_second_point].Y
		dl_x1 := x[size_constraint+2*body.Lines[num_line1].First]
		dl_y1 := x[size_constraint+2*body.Lines[num_line1].First+1]
		dl_x2 := x[size_constraint+2*body.Lines[num_line1].Second]
		dl_y2 := x[size_constraint+2*body.Lines[num_line1].Second+1]
		dl_x3 := x[size_constraint+2*body.Lines[num_line2].First]
		dl_y3 := x[size_constraint+2*body.Lines[num_line2].First+1]
		dl_x4 := x[size_constraint+2*body.Lines[num_line2].Second]
		dl_y4 := x[size_constraint+2*body.Lines[num_line2].Second+1]

		b.SetVec(2*nEqual+2*nFix+nDist+nBelong+nParall+nPerpen+i, (x2+dl_x2-x1-dl_x1)*(x4+dl_x4-x3-dl_x3)+(y2+dl_y2-y1-dl_y1)*(y4+dl_y4-y3-dl_y3))
	}

	for i, value := range body.VerticalLine {
		num_line := value
		line_first_point := body.Lines[num_line].First
		line_second_point := body.Lines[num_line].Second
		x1 := body.Points[line_first_point].X
		x2 := body.Points[line_second_point].X
		dl_x1 := x[size_constraint+2*body.Lines[num_line].First]
		dl_x2 := x[size_constraint+2*body.Lines[num_line].Second]

		b.SetVec(2*nEqual+2*nFix+nDist+nBelong+nParall+nPerpen+nCorner+i, x2+dl_x2-x1-dl_x1)
	}

	for i, value := range body.HorizontLine {
		num_line := value
		line_first_point := body.Lines[num_line].First
		line_second_point := body.Lines[num_line].Second
		y1 := body.Points[line_first_point].Y
		y2 := body.Points[line_second_point].Y
		dl_y1 := x[size_constraint+2*body.Lines[num_line].First+1]
		dl_y2 := x[size_constraint+2*body.Lines[num_line].Second+1]

		b.SetVec(2*nEqual+2*nFix+nDist+nBelong+nParall+nPerpen+nCorner+nVert+i, y2+dl_y2-y1-dl_y1)
	}
	return b
}

func f_jac_creater(body util.BodyHTTP) func(y, x []float64) {
	nEqual := len(body.EqualTwoPoints)
	nFix := len(body.FixationPoint)
	nDist := len(body.DistanceBetweenPoints)
	nBelong := len(body.BelongOfLine)
	nParall := len(body.ParallelTwoLines)
	nPerpen := len(body.PerpenTwoLines)
	nCorner := len(body.CornerTwoLines)
	nVert := len(body.VerticalLine)
	nHoriz := len(body.HorizontLine)

	size_constraint := 2*nEqual + 2*nFix + nDist + nBelong + nParall + nPerpen + nCorner + nVert + nHoriz
	size_matrix := 2*len(body.Points) + size_constraint

	f_jac := func(y, x []float64) {
		// Define your function here

		for i := 0; i < size_matrix; i++ {
			y[i] = 0
		}
		for i := size_constraint; i < size_matrix; i++ {
			y[i] += x[i]
		}

		for i, value := range body.EqualTwoPoints {
			first := value.First
			second := value.Second
			x1 := body.Points[first].X
			y1 := body.Points[first].Y
			x2 := body.Points[second].X
			y2 := body.Points[second].Y
			dl_x1 := x[size_constraint+2*first]
			dl_y1 := x[size_constraint+2*first+1]
			dl_x2 := x[size_constraint+2*second]
			dl_y2 := x[size_constraint+2*second+1]
			l1 := x[2*i]
			l2 := x[2*i+1]

			y[2*i] += x2 + dl_x2 - x1 - dl_x1
			y[2*i+1] += y2 + dl_y2 - y1 - dl_y1
			y[size_constraint+2*first] += -l1
			y[size_constraint+2*first+1] += -l2
			y[size_constraint+2*second] += +l1
			y[size_constraint+2*second+1] += +l2
		}

		for i, value := range body.FixationPoint {
			num_point := value
			x1 := body.Points[num_point].X
			y1 := body.Points[num_point].Y
			xc := x1
			yc := y1
			dl_x1 := x[size_constraint+2*num_point]
			dl_y1 := x[size_constraint+2*num_point+1]
			l1 := x[2*nEqual+2*i]
			l2 := x[2*nEqual+2*i+1]

			y[2*nEqual+2*i] += x1 + dl_x1 - xc
			y[2*nEqual+2*i+1] += y1 + dl_y1 - yc
			y[size_constraint+2*num_point] += l1
			y[size_constraint+2*num_point+1] += l2
		}

		for i, value := range body.DistanceBetweenPoints {
			first := value.First
			second := value.Second
			d := value.Value
			x1 := body.Points[first].X
			y1 := body.Points[first].Y
			x2 := body.Points[second].X
			y2 := body.Points[second].Y
			dl_x1 := x[size_constraint+2*first]
			dl_y1 := x[size_constraint+2*first+1]
			dl_x2 := x[size_constraint+2*second]
			dl_y2 := x[size_constraint+2*second+1]
			l1 := x[2*nEqual+2*nFix+i]

			y[2*nEqual+2*nFix+i] += (x2+dl_x2-x1-dl_x1)*(x2+dl_x2-x1-dl_x1) + (y2+dl_y2-y1-dl_y1)*(y2+dl_y2-y1-dl_y1) - d*d
			y[size_constraint+2*first] += -2 * l1 * (x2 + dl_x2 - x1 - dl_x1)
			y[size_constraint+2*first+1] += -2 * l1 * (y2 + dl_y2 - y1 - dl_y1)
			y[size_constraint+2*second] += +2 * l1 * (x2 + dl_x2 - x1 - dl_x1)
			y[size_constraint+2*second+1] += +2 * l1 * (y2 + dl_y2 - y1 - dl_y1)
		}

		for i, value := range body.BelongOfLine {
			num_point := value.First
			num_line := value.Second
			xp := body.Points[num_point].X
			yp := body.Points[num_point].Y
			x1 := body.Points[body.Lines[num_line].First].X
			y1 := body.Points[body.Lines[num_line].First].Y
			x2 := body.Points[body.Lines[num_line].Second].X
			y2 := body.Points[body.Lines[num_line].Second].Y
			dl_xp := x[size_constraint+2*num_point]
			dl_yp := x[size_constraint+2*num_point+1]
			dl_x1 := x[size_constraint+2*body.Lines[num_line].First]
			dl_y1 := x[size_constraint+2*body.Lines[num_line].First+1]
			dl_x2 := x[size_constraint+2*body.Lines[num_line].Second]
			dl_y2 := x[size_constraint+2*body.Lines[num_line].Second+1]
			l1 := x[2*nEqual+2*nFix+nDist+i]

			y[2*nEqual+2*nFix+nDist+i] += (xp+dl_xp-x1-dl_x1)*(y2+dl_y2-yp-dl_yp) - (x2+dl_x2-xp-dl_xp)*(yp+dl_yp-y1-dl_y1)
			y[size_constraint+2*body.Lines[num_line].First] += -l1 * (y2 + dl_y2 - yp - dl_yp)
			y[size_constraint+2*body.Lines[num_line].First+1] += +l1 * (x2 + dl_x2 - xp - dl_xp)
			y[size_constraint+2*body.Lines[num_line].Second] += -l1 * (yp + dl_yp - y1 - dl_y1)
			y[size_constraint+2*body.Lines[num_line].Second+1] += +l1 * (xp + dl_xp - x1 - dl_x1)
			y[size_constraint+2*num_point] += +l1 * (y2 + dl_y2 - y1 - dl_y1)
			y[size_constraint+2*num_point+1] += +l1 * (x1 + dl_x1 - x2 - dl_x2)
		}

		for i, value := range body.ParallelTwoLines {
			num_line1 := value.First
			num_line2 := value.Second
			line1_first_point := body.Lines[num_line1].First
			line1_second_point := body.Lines[num_line1].Second
			line2_first_point := body.Lines[num_line2].First
			line2_second_point := body.Lines[num_line2].Second
			x1 := body.Points[line1_first_point].X
			y1 := body.Points[line1_first_point].Y
			x2 := body.Points[line1_second_point].X
			y2 := body.Points[line1_second_point].Y
			x3 := body.Points[line2_first_point].X
			y3 := body.Points[line2_first_point].Y
			x4 := body.Points[line2_second_point].X
			y4 := body.Points[line2_second_point].Y
			dl_x1 := x[size_constraint+2*body.Lines[num_line1].First]
			dl_y1 := x[size_constraint+2*body.Lines[num_line1].First+1]
			dl_x2 := x[size_constraint+2*body.Lines[num_line1].Second]
			dl_y2 := x[size_constraint+2*body.Lines[num_line1].Second+1]
			dl_x3 := x[size_constraint+2*body.Lines[num_line2].First]
			dl_y3 := x[size_constraint+2*body.Lines[num_line2].First+1]
			dl_x4 := x[size_constraint+2*body.Lines[num_line2].Second]
			dl_y4 := x[size_constraint+2*body.Lines[num_line2].Second+1]
			l1 := x[2*nEqual+2*nFix+nDist+nBelong+i]

			y[2*nEqual+2*nFix+nDist+nBelong+i] += (x2+dl_x2-x1-dl_x1)*(y4+dl_y4-y3-dl_y3) - (x4+dl_x4-x3-dl_x3)*(y2+dl_y2-y1-dl_y1)
			y[size_constraint+2*body.Lines[num_line1].First] += +l1 * (-1.0 * (y4 + dl_y4 - y3 - dl_y3))
			y[size_constraint+2*body.Lines[num_line1].First+1] += +l1 * (x4 + dl_x4 - x3 - dl_x3)
			y[size_constraint+2*body.Lines[num_line1].Second] += +l1 * (y4 + dl_y4 - y3 - dl_y3)
			y[size_constraint+2*body.Lines[num_line1].Second+1] += +l1 * (-1.0 * (x4 + dl_x4 - x3 - dl_x3))
			y[size_constraint+2*body.Lines[num_line2].First] += +l1 * (y2 + dl_y2 - y1 - dl_y1)
			y[size_constraint+2*body.Lines[num_line2].First+1] += +l1 * (-1.0 * (x2 + dl_x2 - x1 - dl_x1))
			y[size_constraint+2*body.Lines[num_line2].Second] += +l1 * (-1.0 * (y2 + dl_y2 - y1 - dl_y1))
			y[size_constraint+2*body.Lines[num_line2].Second+1] += +l1 * (x2 + dl_x2 - x1 - dl_x1)
		}

		for i, value := range body.PerpenTwoLines {
			num_line1 := value.First
			num_line2 := value.Second
			line1_first_point := body.Lines[num_line1].First
			line1_second_point := body.Lines[num_line1].Second
			line2_first_point := body.Lines[num_line2].First
			line2_second_point := body.Lines[num_line2].Second
			x1 := body.Points[line1_first_point].X
			y1 := body.Points[line1_first_point].Y
			x2 := body.Points[line1_second_point].X
			y2 := body.Points[line1_second_point].Y
			x3 := body.Points[line2_first_point].X
			y3 := body.Points[line2_first_point].Y
			x4 := body.Points[line2_second_point].X
			y4 := body.Points[line2_second_point].Y
			dl_x1 := x[size_constraint+2*body.Lines[num_line1].First]
			dl_y1 := x[size_constraint+2*body.Lines[num_line1].First+1]
			dl_x2 := x[size_constraint+2*body.Lines[num_line1].Second]
			dl_y2 := x[size_constraint+2*body.Lines[num_line1].Second+1]
			dl_x3 := x[size_constraint+2*body.Lines[num_line2].First]
			dl_y3 := x[size_constraint+2*body.Lines[num_line2].First+1]
			dl_x4 := x[size_constraint+2*body.Lines[num_line2].Second]
			dl_y4 := x[size_constraint+2*body.Lines[num_line2].Second+1]
			l1 := x[2*nEqual+2*nFix+nDist+nBelong+nParall+i]

			y[2*nEqual+2*nFix+nDist+nBelong+nParall+i] += (x2+dl_x2-x1-dl_x1)*(x4+dl_x4-x3-dl_x3) + (y2+dl_y2-y1-dl_y1)*(y4+dl_y4-y3-dl_y3)
			y[size_constraint+2*body.Lines[num_line1].First] += +l1 * (-1.0 * (x4 + dl_x4 - x3 - dl_x3))
			y[size_constraint+2*body.Lines[num_line1].First+1] += +l1 * (-1.0 * (y4 + dl_y4 - y3 - dl_y3))
			y[size_constraint+2*body.Lines[num_line1].Second] += +l1 * (x4 + dl_x4 - x3 - dl_x3)
			y[size_constraint+2*body.Lines[num_line1].Second+1] += +l1 * (y4 + dl_y4 - y3 - dl_y3)
			y[size_constraint+2*body.Lines[num_line2].First] += +l1 * (-1.0 * (x2 + dl_x2 - x1 - dl_x1))
			y[size_constraint+2*body.Lines[num_line2].First+1] += +l1 * (-1.0 * (y2 + dl_y2 - y1 - dl_y1))
			y[size_constraint+2*body.Lines[num_line2].Second] += +l1 * (x2 + dl_x2 - x1 - dl_x1)
			y[size_constraint+2*body.Lines[num_line2].Second+1] += +l1 * (y2 + dl_y2 - y1 - dl_y1)
		}

		for i, value := range body.CornerTwoLines {
			num_line1 := value.First
			num_line2 := value.Second
			line1_first_point := body.Lines[num_line1].First
			line1_second_point := body.Lines[num_line1].Second
			line2_first_point := body.Lines[num_line2].First
			line2_second_point := body.Lines[num_line2].Second
			x1 := body.Points[line1_first_point].X
			y1 := body.Points[line1_first_point].Y
			x2 := body.Points[line1_second_point].X
			y2 := body.Points[line1_second_point].Y
			x3 := body.Points[line2_first_point].X
			y3 := body.Points[line2_first_point].Y
			x4 := body.Points[line2_second_point].X
			y4 := body.Points[line2_second_point].Y
			dl_x1 := x[size_constraint+2*body.Lines[num_line1].First]
			dl_y1 := x[size_constraint+2*body.Lines[num_line1].First+1]
			dl_x2 := x[size_constraint+2*body.Lines[num_line1].Second]
			dl_y2 := x[size_constraint+2*body.Lines[num_line1].Second+1]
			dl_x3 := x[size_constraint+2*body.Lines[num_line2].First]
			dl_y3 := x[size_constraint+2*body.Lines[num_line2].First+1]
			dl_x4 := x[size_constraint+2*body.Lines[num_line2].Second]
			dl_y4 := x[size_constraint+2*body.Lines[num_line2].Second+1]
			l1 := x[2*nEqual+2*nFix+nDist+nBelong+nParall+nPerpen+i]

			y[2*nEqual+2*nFix+nDist+nBelong+nParall+nPerpen+i] += (x2+dl_x2-x1-dl_x1)*(x4+dl_x4-x3-dl_x3) + (y2+dl_y2-y1-dl_y1)*(y4+dl_y4-y3-dl_y3)
			y[size_constraint+2*body.Lines[num_line1].First] += +l1 * (-1.0 * (x4 + dl_x4 - x3 - dl_x3))
			y[size_constraint+2*body.Lines[num_line1].First+1] += +l1 * (-1.0 * (y4 + dl_y4 - y3 - dl_y3))
			y[size_constraint+2*body.Lines[num_line1].Second] += +l1 * (x4 + dl_x4 - x3 - dl_x3)
			y[size_constraint+2*body.Lines[num_line1].Second+1] += +l1 * (y4 + dl_y4 - y3 - dl_y3)
			y[size_constraint+2*body.Lines[num_line2].First] += +l1 * (-1.0 * (x2 + dl_x2 - x1 - dl_x1))
			y[size_constraint+2*body.Lines[num_line2].First+1] += +l1 * (-1.0 * (y2 + dl_y2 - y1 - dl_y1))
			y[size_constraint+2*body.Lines[num_line2].Second] += +l1 * (x2 + dl_x2 - x1 - dl_x1)
			y[size_constraint+2*body.Lines[num_line2].Second+1] += +l1 * (y2 + dl_y2 - y1 - dl_y1)
		}

		for i, value := range body.VerticalLine {
			num_line := value
			line_first_point := body.Lines[num_line].First
			line_second_point := body.Lines[num_line].Second
			x1 := body.Points[line_first_point].X
			x2 := body.Points[line_second_point].X
			dl_x1 := x[size_constraint+2*body.Lines[num_line].First]
			dl_x2 := x[size_constraint+2*body.Lines[num_line].Second]
			l1 := x[2*nEqual+2*nFix+nDist+nBelong+nParall+nPerpen+nCorner+i]

			y[2*nEqual+2*nFix+nDist+nBelong+nParall+nPerpen+nCorner+i] += x2 + dl_x2 - x1 - dl_x1
			y[size_constraint+2*body.Lines[num_line].First] += -l1
			y[size_constraint+2*body.Lines[num_line].First+1] += 0
			y[size_constraint+2*body.Lines[num_line].Second] += +l1
			y[size_constraint+2*body.Lines[num_line].Second+1] += 0
		}

		for i, value := range body.HorizontLine {
			num_line := value
			line_first_point := body.Lines[num_line].First
			line_second_point := body.Lines[num_line].Second
			y1 := body.Points[line_first_point].Y
			y2 := body.Points[line_second_point].Y
			dl_y1 := x[size_constraint+2*body.Lines[num_line].First+1]
			dl_y2 := x[size_constraint+2*body.Lines[num_line].Second+1]
			l1 := x[2*nEqual+2*nFix+nDist+nBelong+nParall+nPerpen+nCorner+nVert+i]

			y[2*nEqual+2*nFix+nDist+nBelong+nParall+nPerpen+nCorner+nVert+i] += y2 + dl_y2 - y1 - dl_y1
			y[size_constraint+2*body.Lines[num_line].First] += 0
			y[size_constraint+2*body.Lines[num_line].First+1] += -l1
			y[size_constraint+2*body.Lines[num_line].Second] += 0
			y[size_constraint+2*body.Lines[num_line].Second+1] += +l1
		}
	}
	return f_jac
}

// Метод Ньютона для решения СЛАУ.
func newtonMethod(body util.BodyHTTP) (util.BodyHTTP, error) {

	nEqual := len(body.EqualTwoPoints)
	nFix := len(body.FixationPoint)
	nDist := len(body.DistanceBetweenPoints)
	nBelong := len(body.BelongOfLine)
	nParall := len(body.ParallelTwoLines)
	nPerpen := len(body.PerpenTwoLines)
	nCorner := len(body.CornerTwoLines)
	nVert := len(body.VerticalLine)
	nHoriz := len(body.HorizontLine)
	size_constraint := 2*nEqual + 2*nFix + nDist + nBelong + nParall + nPerpen + nCorner + nVert + nHoriz
	size_matrix := 2*len(body.Points) + size_constraint

	f_jac := f_jac_creater(body)

	const maxIterations = 1000
	const epsilon = 1e-8

	// Определите начальное приближение x0.
	x := mat.NewVecDense(size_matrix, make([]float64, size_matrix))
	x_st := make([]float64, size_matrix)
	for i := 0; i < size_matrix; i++ {
		x_st[i] = 0
		x.SetVec(i, 0)
	}

	for k := 0; k < maxIterations; k++ {
		// Create a matrix
		jac := mat.NewDense(size_matrix, size_matrix, nil)
		fd.Jacobian(jac, f_jac, x_st, &fd.JacobianSettings{
			Formula:    fd.Central,
			Concurrent: true,
		})
		// Create a vector b
		b := b_creater(body, x_st)

		// Проверьте условие сходимости.
		maxDiff := 0.0
		for i := 0; i < size_matrix; i++ {
			if math.Abs(b.AtVec(i)) > maxDiff { // f // attention
				maxDiff = math.Abs(b.AtVec(i))
			}
		}

		if maxDiff < epsilon {
			for i, _ := range body.Points {
				body.Points[i].X += x.AtVec(size_constraint + 2*i)
				body.Points[i].Y += x.AtVec(size_constraint + 2*i + 1)
			}
			return body, nil
		}

		b.ScaleVec(-1.0, b)
		// Решите систему линейных уравнений J * dx = -b для приращения dx.
		dx, err := solveLinearSystem(jac, b)
		if err != nil {
			fmt.Printf("Error after solve")
			return body, err
		}

		// Обновите текущее приближение x.
		for i := 0; i < size_matrix; i++ {
			x.SetVec(i, x.AtVec(i)+dx.AtVec(i))
			x_st[i] = x_st[i] + dx.AtVec(i)
		}
		//fmt.Println("x: ", x)
	}
	return body, &errPkg.MyErrors{ProjectTypeText: errPkg.NotFoundSolver, SourceText: errPkg.NotFoundSolver, Way: "newtonMethod"}
}

func solveLinearSystem(A *mat.Dense, b *mat.VecDense) (*mat.VecDense, error) {
	//for i := 0; i < b.Len(); i++ {
	//	for j := 0; j < b.Len(); j++ {
	//		fmt.Printf("  %.2f", A.At(i, j))
	//	}
	//	fmt.Printf("\n")
	//}
	//fmt.Printf("\nb: ", b)
	var x mat.VecDense
	errSolver := x.SolveVec(A, b)
	if errSolver != nil {
		return nil, &errPkg.MyErrors{ProjectTypeText: errPkg.NotFoundSolver, SourceText: errSolver.Error(), Way: "solveLinearSystem"}
	}
	return &x, nil
}
