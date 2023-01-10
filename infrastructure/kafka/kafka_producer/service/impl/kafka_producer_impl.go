package impl

import (
	"infrastructure/kafka/kafka_producer/exception"
	"net/url"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	kafkaAvro "github.com/mycujoo/go-kafka-avro/v2"
)

type KafkaProducerImpl[K, V any] struct {
	*kafkaAvro.Producer
}

func NewKafkaProducerImpl[K, V any](config kafka.ConfigMap, schemaRegistryUrl url.URL) *KafkaProducerImpl[K, V] {
	p, err := kafkaAvro.NewProducer("", "", "",
		kafkaAvro.WithKafkaConfig(&config),
		kafkaAvro.WithSchemaRegistryURL(&schemaRegistryUrl),
	)
	if err != nil {
		panic(err)
	}
	return &KafkaProducerImpl[K, V]{
		Producer: p,
	}
}

func (k KafkaProducerImpl[K, V]) Send(key K, message V) error {
	err := k.Produce(key, message, nil)
	defer k.Close()
	if err != nil {
		return exception.KafkaProducerException("Error while sending message to kafka topic")
	}
	return nil
}
