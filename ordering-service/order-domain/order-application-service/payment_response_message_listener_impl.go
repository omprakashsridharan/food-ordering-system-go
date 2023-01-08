package order_application_service

import "order-application-service/dto/message"

type PaymentResponseMessageListenerImpl struct {
}

func (p *PaymentResponseMessageListenerImpl) PaymentCompleted(response message.PaymentResponse) {
	//TODO implement me
	panic("implement me")
}

func (p *PaymentResponseMessageListenerImpl) PaymentCancelled(response message.PaymentResponse) {
	//TODO implement me
	panic("implement me")
}

func NewPaymentResponseMessageListenerImpl() *PaymentResponseMessageListenerImpl {
	return &PaymentResponseMessageListenerImpl{}
}
