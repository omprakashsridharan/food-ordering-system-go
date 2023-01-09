package exception

import "errors"

func KafkaProducerException(message string) error {
	return errors.New(message)
}
