package apiControllerV1

import (
	keycloak "antriin/src/modules/keycloak/auth"
	"github.com/labstack/echo/v4"
)

func AuthController(e *echo.Echo, kc keycloak.IKeycloak)  {
	handler := NewAuthHandler(kc)

	endpointGroup := e.Group("v1/auth")
	endpointGroup.POST("/login", handler.UserLogin)
}