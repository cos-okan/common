package common

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
