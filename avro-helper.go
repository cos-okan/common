package common

import (
	"embed"
	"log"

	"github.com/hamba/avro/v2"
)

var (
	twrDistanceAvroSchema       avro.Schema
	mdUpdateAvroSchema          avro.Schema
	processedDistanceAvroSchema avro.Schema
)

//go:embed avsc/*
var avsc embed.FS

func PrepareAvroHelper() {
	twrDistanceAvroSchema, _ = prepareAvroSchema("avsc/twr-distance.avsc")
	mdUpdateAvroSchema, _ = prepareAvroSchema("avsc/md-update.avsc")
	processedDistanceAvroSchema, _ = prepareAvroSchema("avsc/processed-distance.avsc")
}

func prepareAvroSchema(avroSchemaFilePath string) (schema avro.Schema, err error) {
	avroSchemaByte, err := avsc.ReadFile(avroSchemaFilePath)
	if err != nil {
		log.Fatal(err)
	}

	schema, err = avro.Parse(string(avroSchemaByte))
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (td *TwrDistance) AvroSerializer() (data []byte, err error) {
	data, err = avro.Marshal(twrDistanceAvroSchema, td)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (mdu *MasterDataUpdate) AvroSerializer() (data []byte, err error) {
	data, err = avro.Marshal(mdUpdateAvroSchema, mdu)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (pd *ProcessedDistance) AvroSerializer() (data []byte, err error) {
	data, err = avro.Marshal(processedDistanceAvroSchema, pd)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (td *TwrDistance) AvroDeserializer(data []byte) (err error) {
	return avro.Unmarshal(twrDistanceAvroSchema, data, &td)
}

func (mdu *MasterDataUpdate) AvroDeserializer(data []byte) (err error) {
	return avro.Unmarshal(mdUpdateAvroSchema, data, &mdu)
}

func (pd *ProcessedDistance) AvroDeserializer(data []byte) (err error) {
	return avro.Unmarshal(processedDistanceAvroSchema, data, &pd)
}
