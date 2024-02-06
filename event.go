package common

import "time"

type RedpandaEvent interface {
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
	IsUndefined        bool        `avro:"isUndefined"`
	IsOutOfRange       bool        `avro:"isOutOfRange"`
	IsShort            bool        `avro:"isShort"`
	IsLong             bool        `avro:"isLong"`
	OnAnchor           bool        `avro:"onAnchor"`
	ConfidenceLevel    int         `avro:"confidenceLevel"`
}
