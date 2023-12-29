package util

type BodyHTTP struct { //TODO(N): remake
	Points                []Point            `json:"points"`                // Коордианты точек. Последовательность важна, влияет ниже
	Lines                 []PairNumber       `json:"lines"`                 // Номера точек
	EqualTwoPoints        []PairNumber       `json:"equalTwoPoints"`        // Номера точек
	DistanceBetweenPoints []DistancePoints   `json:"distanceBetweenPoints"` // Номера точек
	FixationPoint         []int              `json:"fixationPoint"`         // Массив из номеров точек, которые зафиксированы
	BelongOfLine          []PairNumber       `json:"belongOfLine"`          // Первый - номер точки, второй - номер линии
	ParallelTwoLines      []PairNumber       `json:"parallelTwoLines"`      // Номера линий
	PerpenTwoLines        []PairNumber       `json:"perpenTwoLines"`        // Номера линий
	CornerTwoLines        []PairNumberCorner `json:"cornerTwoLines"`        // Номера линий
	VerticalLine          []int              `json:"verticalLine"`          // Массив из номеров линий
	HorizontLine          []int              `json:"horizontLine"`          // Массив из номеров линий
}

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type PairNumber struct {
	First  int `json:"first"`
	Second int `json:"second"`
}

type PairNumberCorner struct {
	First  int `json:"first"`
	Second int `json:"second"`
	Value  int `json:"value"`
}

type DistancePoints struct {
	Value  float64 `json:"value"`
	First  int     `json:"first"`
	Second int     `json:"second"`
}
