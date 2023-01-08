package order_application_service

import "order-application-service/dto/message"

type RestaurantApprovalResponseMessageListenerImpl struct {
}

func (r *RestaurantApprovalResponseMessageListenerImpl) OrderApproved(response message.RestaurantApprovalResponse) {
	//TODO implement me
	panic("implement me")
}

func (r *RestaurantApprovalResponseMessageListenerImpl) OrderRejected(response message.RestaurantApprovalResponse) {
	//TODO implement me
	panic("implement me")
}

func NewRestaurantApprovalResponseMessageListenerImpl() *RestaurantApprovalResponseMessageListenerImpl {
	return &RestaurantApprovalResponseMessageListenerImpl{}
}
