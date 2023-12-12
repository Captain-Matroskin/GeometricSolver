package util

type RequestBody struct { //TODO(N): remake
	Points []Point `json:"points"`
	Lines  []Line  `json:"lines"`
}

type Point struct {
	x float64
	y float64
}

type Line struct {
	Point1 Point
	Point2 Point
}

type TwoPoints struct {
	Point1 Point
	Point2 Point
}
