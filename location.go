package common

import "math"

const (
	TwoCirclePerfectIntersectionDistance = 100
	StationaryDistanceChangeThreshold    = 50
)

type Position struct {
	EntityID        int      `json:"EntityID"`
	Location        Location `json:"Location"`
	Date            string   `json:"Date"`
	ConfidenceLevel int      `json:"ConfidenceLevel"`
	Speed           int      `json:"Speed"`
	Direction       int      `json:"Direction"`
	EstimationRange int      `json:"EstimationRange"`
}

type Point struct {
	X int `avro:"x"`
	Y int `avro:"y"`
	Z int `avro:"z"`
}

type Location struct {
	FloorID int   `avro:"floorId"`
	Point   Point `avro:"point"`
}

type Circle struct {
	C Point
	R int
}

type CircleIntersectionResult int

const (
	NoIntersection CircleIntersectionResult = iota
	HasIntersection
	PerfectIntersection
)

func FindCircleIntersectionPoints(c1 Circle, c2 Circle) (intersections [2]Point, result CircleIntersectionResult) {
	d := CalculateDistance(c1.C, c2.C)
	if d > float64(c1.R+c2.R) { // TODO : fark çok küçükse çemberleri çok az büyüt
		result = NoIntersection
		return
	}

	if d < math.Abs(float64(c1.R-c2.R)) { // TODO : fark çok küçükse, büyük çemberi çok az küçült
		result = NoIntersection
		return
	}

	a := (math.Pow(float64(c1.R), 2) - math.Pow(float64(c2.R), 2) + math.Pow(d, 2)) / (2 * d)
	h := math.Sqrt(math.Pow(float64(c1.R), 2) - math.Pow(a, 2))
	cx := float64(c1.C.X) + (a*(float64(c2.C.X-c1.C.X)))/d
	cy := float64(c1.C.Y) + (a*(float64(c2.C.Y-c1.C.Y)))/d

	intersections[0] = Point{X: int(cx + (h*(float64(c2.C.Y-c1.C.Y)))/d), Y: int(cy - (h*(float64(c2.C.X-c1.C.X)))/d)}
	intersections[1] = Point{X: int(cx - (h*(float64(c2.C.Y-c1.C.Y)))/d), Y: int(cy + (h*(float64(c2.C.X-c1.C.X)))/d)}

	if h*2 < TwoCirclePerfectIntersectionDistance {
		result = PerfectIntersection
	} else {
		result = HasIntersection
	}

	return
}

func CalculateDistance(p1 Point, p2 Point) float64 {
	return math.Sqrt(math.Pow(float64(p1.X-p2.X), 2)+math.Pow(float64(p1.Y-p2.Y), 2)) + math.Pow(float64(p1.Z-p2.Z), 2)
}
