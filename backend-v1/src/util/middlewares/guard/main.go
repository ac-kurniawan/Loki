package guard

import (
	keycloak "antriin/src/modules/keycloak/auth"
	"github.com/labstack/echo/v4"
	"strings"
)

type GuardErrorResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

func GuardMiddleware(kc keycloak.IKeycloak) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// get bearer token
			token := c.Request().Header.Get("Authorization")
			if token == "" {
				return c.JSON(401, GuardErrorResponse{
					Code: 401,
					Message: "token doesn't exist",
				})
			}

			// Get token value
			splitToken := strings.Split(token, "Bearer ")
			token = splitToken[1]

			// validate token
			valid, err := kc.ValidateToken(token)
			if err != nil || !*valid.Active {
				return c.JSON(401, GuardErrorResponse{
					Code: 401,
					Message: "token doesn't valid",
				})
			}

			// store token in user context
			claims, err := kc.DecodeToken(token)
			if err != nil {
				return c.JSON(401, GuardErrorResponse{
					Code: 401,
					Message: "token doesn't valid",
				})
			}
			claimDecode := claims.Claims.(*keycloak.KeycloakTokenClaims)
			c.Set("userId", claimDecode.Sub)
			c.Set("userEmail", claimDecode.Email)

			return next(c)
		}
	}
}
