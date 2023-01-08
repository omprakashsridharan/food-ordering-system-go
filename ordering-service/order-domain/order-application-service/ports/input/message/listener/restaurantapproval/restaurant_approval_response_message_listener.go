package restaurantapproval

import "order-application-service/dto/message"

type RestaurantApprovalResponseMessageListener interface {
	OrderApproved(response message.RestaurantApprovalResponse)
	OrderRejected(response message.RestaurantApprovalResponse)
}
