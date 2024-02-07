package common

import (
	"context"

	"github.com/twmb/franz-go/pkg/kgo"
)

type RedpandaProducer struct {
	client *kgo.Client
	topic  string
}

type RedpandaConsumer struct {
	Client *kgo.Client
	topic  string
}

func NewRedpandaProducer(brokers []string, topic string) *RedpandaProducer {
	PrepareAvroHelper()
	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
	)
	if err != nil {
		panic(err)
	}
	return &RedpandaProducer{client: client, topic: topic}
}

func NewRedpandaConsumer(brokers []string, topic string, twrConsumerGroupID string) *RedpandaConsumer {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.ConsumerGroup(twrConsumerGroupID),
		kgo.ConsumeTopics(topic),
		kgo.ConsumeResetOffset(kgo.NewOffset().AtStart()),
	)
	if err != nil {
		panic(err)
	}
	return &RedpandaConsumer{Client: client, topic: topic}
}

func (p *RedpandaProducer) SendAvroMessage(rpm AvroEvent, key []byte) {
	ctx := context.Background()

	serializedData, err := rpm.AvroSerializer()

	if err != nil {
		return
	}

	record := kgo.Record{
		Topic: p.topic,
		Value: serializedData,
	}

	if key != nil {
		record.Key = key
	}

	p.client.Produce(ctx, &record, nil)
}
