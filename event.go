package common

import (
	"log"
	"time"

	"github.com/hamba/avro/v2"
)

type AvroEvent interface {
	AvroSerializer() (data []byte, err error)
	AvroDeserializer(data []byte) (err error)
}

type TwrDistance struct {
	FromNodeId        int       `avro:"fromNodeId"`
	ToNodeId          int       `avro:"toNodeId"`
	MessageNo         int       `avro:"messageNo"`
	Distance          int       `avro:"distance"`
	FwConfidenceLevel int       `avro:"fwConfidenceLevel"`
	Timestamp         time.Time `avro:"timestamp"`
}

type ProcessedDistance struct {
	Twr                TwrDistance `avro:"twr"`
	MessageType        int         `avro:"messageType"`
	Entity             Entity      `avro:"entity"`
	Anchor             Anchor      `avro:"anchor"`
	ProjectionDistance int         `avro:"projectionDistance"`
	IsInvalid          bool        `avro:"isInvalid"`
	InvalidReason      int         `avro:"invalidReason"`
	OutOfRange         bool        `avro:"outOfRange"`
	OnAnchor           bool        `avro:"onAnchor"`
	ConfidenceLevel    int         `avro:"confidenceLevel"`
	Timestamp          time.Time   `avro:"timestamp"`
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

type InvalidReason int

const (
	Undefined InvalidReason = iota + 1
	Short
	Long
	SudokuConflict
	SuspiciousTwrChange
	IncompatibleWithEstimatedPosition
	IncompatibleWithOtherAnchorMeasurement
)

type ConfidenceLevel int

const (
	LowestConfidence ConfidenceLevel = iota + 1
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
