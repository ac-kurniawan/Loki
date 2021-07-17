package apiControllerV1

import (
	"antriin/src/business/attendee"
	keycloak "antriin/src/modules/keycloak/auth"
	"antriin/src/util/middlewares/guard"
	"github.com/labstack/echo/v4"
)

func AttendeeController(echo *echo.Echo, attendee attendee.IAttendee,
	kc keycloak.IKeycloak)  {
	handler := NewAttendeeHandler(attendee)

	endpointGroup := echo.Group("v1/attendee")
	endpointGroup.GET("/list", handler.getAttendees, guard.GuardMiddleware(kc))
	endpointGroup.POST("", handler.setAttendee, guard.GuardMiddleware(kc))

}