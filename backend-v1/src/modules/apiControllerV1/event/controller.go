package apiControllerV1

import (
	"antriin/src/business/event"
	keycloak "antriin/src/modules/keycloak/auth"
	"antriin/src/util/middlewares/guard"
	"github.com/labstack/echo/v4"
)

func EventController(echo *echo.Echo, event event.IEvent,
	kc keycloak.IKeycloak)  {
	handler := NewEventHandler(event)

	endpointGroup := echo.Group("v1/event")
	endpointGroup.POST("", handler.SetEvent, guard.GuardMiddleware(kc))
	endpointGroup.GET("", handler.GetEventsByCreatorID,
		guard.GuardMiddleware(kc))
	endpointGroup.POST("/:eventId/schedule", handler.CreateScheduleByEvent,
		guard.GuardMiddleware(kc))

}