package util

type BodyHTTP struct { //TODO(N): remake
	Points                []Point          `json:"points"`                // Коордианты точек. Последовательность важна, влияет ниже
	Lines                 []PairNumber     `json:"lines"`                 // Номера точек
	EqualTwoPoints        []PairNumber     `json:"equalTwoPoints"`        // Номера точек
	DistanceBetweenPoints []DistancePoints `json:"distanceBetweenPoints"` // Номера точек
	FixationPoint         []int            // Массив из номеров точек, которые зафиксированы
	BelongOfLine          []PairNumber     // Первый - номер точки, второй - номер линии
	ParallelTwoLines      []PairNumber     `json:"parallelTwoLines"` // Номера линий
	PerpenTwoLines        []PairNumber     // Номера линий
	CornerTwoLines        []PairNumber     // Номера линий
	VerticalLine          []int            // Массив из номеров линий
	HorizontLine          []int            // Массив из номеров линий
}

type Point struct {
	X float64
	Y float64
}

type PairNumber struct {
	First  int
	Second int
}

type DistancePoints struct {
	Distance float64
	Points   PairNumber
}
