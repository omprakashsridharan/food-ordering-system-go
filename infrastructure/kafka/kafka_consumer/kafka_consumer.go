package kafkaconsumer

type KafkaConsumer[T any] interface {
	Receive(messages []T)
}
