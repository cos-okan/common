package common

type MasterData struct {
	Anchors  map[string]Anchor
	Tags     map[string]int
	Entities map[string]Entity
}

type Anchor struct {
	ID       int      `avro:"id"`
	Location Location `avro:"location"`
	Range    int      `avro:"range"`
	Sudoku   int      `avro:"sudoku"`
}

type Entity struct {
	ID       int `avro:"id"`
	Height   int `avro:"height"`
	TagID    int `avro:"tagId"`
	MaxSpeed int `avro:"maxSpeed"`
	Type     int `avro:"type"`
}
