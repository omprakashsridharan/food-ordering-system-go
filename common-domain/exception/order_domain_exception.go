package exception

func OrderDomainException(message string) error {
	return DomainException(message)
}
