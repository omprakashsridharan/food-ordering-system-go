package exception

func OrderNotFoundException(message string) error {
	return DomainException(message)
}
