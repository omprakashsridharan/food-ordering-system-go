package service

type KafkaProducer[K any, V any] interface {
	Send(key K, message V) error
}
