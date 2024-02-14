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
	twrGroupAvroSchema          avro.Schema
)

//go:embed avsc/*
var avsc embed.FS

func PrepareAvroHelper() {
	twrDistanceAvroSchema, _ = prepareAvroSchema("avsc/twr-distance.avsc")
	mdUpdateAvroSchema, _ = prepareAvroSchema("avsc/md-update.avsc")
	processedDistanceAvroSchema, _ = prepareAvroSchema("avsc/processed-distance.avsc")
	// twrGroupAvroSchema, _ = prepareAvroSchema("avsc/twr-group.avsc") // TODO: çalışr hale getirelecek
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
