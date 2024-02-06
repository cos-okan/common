package common

import (
	"time"
)

type MasterData struct {
	Anchors  map[string]Anchor
	Tags     map[string]int
	Entities map[string]Entity
}

type Anchor struct {
	ID       int      `avro:"id"`
	Location Location `avro:"location"`
	Range    int      `avro:"range"`
}

type Entity struct {
	ID       int `avro:"id"`
	Height   int `avro:"height"`
	TagID    int `avro:"tagId"`
	MaxSpeed int `avro:"maxSpeed"`
	Type     int `avro:"type"`
}

type MasterDataUpdate struct {
	Operation int       `avro:"operation"`
	DataType  int       `avro:"dataType"`
	Key       string    `avro:"key"`
	Anchor    Anchor    `avro:"anchor"`
	Tag       int       `avro:"tag"`
	Entity    Entity    `avro:"entity"`
	Timestamp time.Time `avro:"timestamp"`
}
