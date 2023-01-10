package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"order-application-service/dto/create"
	"order-application-service/dto/track"
	oas "order-application-service/ports/input/service"
)

func CreateOrder(orderApplicationService oas.OrderApplicationService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		createOrderCommand := new(create.CreateOrderCommand)
		if err := c.BodyParser(createOrderCommand); err != nil {
			return err
		}
		createOrderResponse, createOrderErr := orderApplicationService.CreateOrder(*createOrderCommand)
		if createOrderErr != nil {
			return createOrderErr
		}
		return c.JSON(createOrderResponse)
	}
}

func GetOrderByTrackingId(orderApplicationService oas.OrderApplicationService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		trackingIdUuid, err := uuid.Parse(c.Params("trackingId"))
		if err != nil {
			return err
		}
		trackOrderQuery := track.NewTrackOrderQuery(trackingIdUuid)
		trackOrderResponse, trackOrderErr := orderApplicationService.TrackOrder(*trackOrderQuery)
		if trackOrderErr != nil {
			return trackOrderErr
		}
		return c.JSON(trackOrderResponse)
	}
}
