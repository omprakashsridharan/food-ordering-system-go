package payment

import "order-application-service/dto/message"

type PaymentResponseListener interface {
	PaymentCompleted(response message.PaymentResponse)
	PaymentCancelled(response message.PaymentResponse)
}
