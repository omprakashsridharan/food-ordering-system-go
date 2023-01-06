package exception

import "errors"

func DomainException(message string) error {
	return errors.New(message)
}
