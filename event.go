package common

import (
	"encoding/json"
	"log"
	"time"

	"github.com/hamba/avro/v2"
)

type Event interface {
	AvroSerializer() (data []byte, err error)
	AvroDeserializer(data []byte) (err error)
	JsonSerializer() (data []byte, err error)
	JsonDeserializer(data []byte) (err error)
}

type TwrDistance struct {
	FromNodeId        int       `avro:"fromNodeId" json:"fromNodeId"`
	ToNodeId          int       `avro:"toNodeId" json:"toNodeId"`
	MessageNo         int       `avro:"messageNo" json:"messageNo"`
	Distance          int       `avro:"distance" json:"distance"`
	FwConfidenceLevel int       `avro:"fwConfidenceLevel" json:"fwConfidenceLevel"`
	Timestamp         time.Time `avro:"timestamp" json:"timestamp"`
}

type ProcessedDistance struct {
	Twr                TwrDistance `avro:"twr" json:"twr"`
	MessageType        int         `avro:"messageType" json:"messageType"`
	Entity             Entity      `avro:"entity" json:"entity"`
	Anchor             Anchor      `avro:"anchor" json:"anchor"`
	ProjectionDistance int         `avro:"projectionDistance" json:"projectionDistance"`
	IsInvalid          bool        `avro:"isInvalid" json:"isInvalid"`
	InvalidReason      int         `avro:"invalidReason" json:"invalidReason"`
	OutOfRange         bool        `avro:"outOfRange" json:"outOfRange"`
	OnAnchor           bool        `avro:"onAnchor" json:"onAnchor"`
	ConfidenceLevel    int         `avro:"confidenceLevel" json:"confidenceLevel"`
	Timestamp          time.Time   `avro:"timestamp" json:"timestamp"`
}

type MasterDataUpdate struct {
	Operation int       `avro:"operation" json:"operation"`
	DataType  int       `avro:"dataType" json:"dataType"`
	Key       string    `avro:"key" json:"key"`
	Anchor    Anchor    `avro:"anchor" json:"anchor"`
	Tag       int       `avro:"tag" json:"tag"`
	Entity    Entity    `avro:"entity" json:"entity"`
	Timestamp time.Time `avro:"timestamp" json:"timestamp"`
}

type TwrGroup struct {
	EntityID        int                        `avro:"entityID"  json:"entityID"`
	TwrCycleNo      int                        `avro:"twrCycleNo" json:"twrCycleNo"`
	FloorID         int                        `avro:"floorID" json:"floorID"`
	Timestamp       time.Time                  `avro:"timestamp" json:"timestamp"`
	IsStationary    bool                       `avro:"isStationary" json:"isStationary"`
	AnchorEventMap  map[int]ProcessedDistance  `avro:"anchorEventMap" json:"anchorEventMap"`
	IntersectionMap map[int](map[int][2]Point) `avro:"intersectionMap" json:"intersectionMap"`
}

type InvalidReason int

const (
	Undefined InvalidReason = iota + 1
	Short
	Long
	SudokuConflict
	DifferentFloor
	IncompatibleWithEstimatedPosition
	SuspiciousTwrChange
	StationaryControlFailure
	IncompatibleWithOtherAnchorMeasurement
)

type TwrType int

const (
	TagToAnchor TwrType = iota + 1
	AnchorToAnchor
	TagToTag
)

type ConfidenceLevel int

const (
	NoConfidence ConfidenceLevel = iota
	LowestConfidence
	LowConfidence
	MediumConfidence
	HighConfidence
	HighestConfidence
)

type AccuracyLevel int

const (
	LowestAccuracy AccuracyLevel = iota + 1
	LowAccuracy
	MediumAccuracy
	HighAccuracy
	HighestAccuracy
)

func (td *TwrDistance) AvroSerializer() (data []byte, err error) {
	data, err = avro.Marshal(twrDistanceAvroSchema, td)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (td *TwrDistance) AvroDeserializer(data []byte) (err error) {
	return avro.Unmarshal(twrDistanceAvroSchema, data, &td)
}

func (td *TwrDistance) JsonSerializer() (data []byte, err error) {
	data, err = json.Marshal(td)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (td *TwrDistance) JsonDeserializer(data []byte) (err error) {
	return json.Unmarshal(data, &td)
}

func (pd *ProcessedDistance) AvroSerializer() (data []byte, err error) {
	data, err = avro.Marshal(processedDistanceAvroSchema, pd)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (pd *ProcessedDistance) AvroDeserializer(data []byte) (err error) {
	return avro.Unmarshal(processedDistanceAvroSchema, data, &pd)
}

func (pd *ProcessedDistance) JsonSerializer() (data []byte, err error) {
	data, err = json.Marshal(pd)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (pd *ProcessedDistance) JsonDeserializer(data []byte) (err error) {
	return json.Unmarshal(data, &pd)
}

func (mdu *MasterDataUpdate) AvroSerializer() (data []byte, err error) {
	data, err = avro.Marshal(mdUpdateAvroSchema, mdu)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (mdu *MasterDataUpdate) AvroDeserializer(data []byte) (err error) {
	return avro.Unmarshal(mdUpdateAvroSchema, data, &mdu)
}

func (mdu *MasterDataUpdate) JsonSerializer() (data []byte, err error) {
	data, err = json.Marshal(mdu)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (mdu *MasterDataUpdate) JsonDeserializer(data []byte) (err error) {
	return json.Unmarshal(data, &mdu)
}

func (twrGroup *TwrGroup) AvroSerializer() (data []byte, err error) {
	data, err = avro.Marshal(twrGroupAvroSchema, twrGroup)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (twrGroup *TwrGroup) AvroDeserializer(data []byte) (err error) {
	return avro.Unmarshal(twrGroupAvroSchema, data, &twrGroup)
}

func (twrGroup *TwrGroup) JsonSerializer() (data []byte, err error) {
	data, err = json.Marshal(twrGroup)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (twrGroup *TwrGroup) JsonDeserializer(data []byte) (err error) {
	return json.Unmarshal(data, &twrGroup)
}

func MakeInvalid(p *ProcessedDistance, invalidReason InvalidReason) {
	p.IsInvalid = true
	p.InvalidReason = int(invalidReason)
	p.ConfidenceLevel = int(NoConfidence)
}
