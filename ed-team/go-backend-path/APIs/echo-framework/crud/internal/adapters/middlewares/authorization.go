package middlewares

import (
	"github.com/labstack/echo"

	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/adapters/handlers"
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/usecase/authorize/tokens"
)

func Authorization(f handlers.FunctionType) handlers.FunctionType {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		_, err := tokens.ValidateToken(token)
		if err != nil {
			return handlers.Forbidden(c)
		}

		return f(c)
	}
}
