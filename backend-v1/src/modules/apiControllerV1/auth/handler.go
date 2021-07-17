package apiControllerV1

import (
	keycloak "antriin/src/modules/keycloak/auth"
	"antriin/src/util/response"
	"github.com/labstack/echo/v4"
)

type IAuthHandler interface {
	UserLogin(ctx echo.Context) error
}

type Handler struct {
	keycloakModule keycloak.IKeycloak
}

func (h Handler) UserLogin(ctx echo.Context) error {
	req := new(SetAuthLoginRequest)
	if err := ctx.Bind(&req); err != nil {
		msg := err.Error()
		return ctx.JSON(400, response.Error(400, nil, &msg))
	}

	token, err := h.keycloakModule.Login(req.Username, req.Password)
	if err != nil {
		msg := err.Error()
		return ctx.JSON(401, response.Error(401, nil, &msg))
	}

	return ctx.JSON(200, response.Success(200, token))
}

func NewAuthHandler(keycloakModule keycloak.IKeycloak) IAuthHandler {
	return Handler{keycloakModule}
}
