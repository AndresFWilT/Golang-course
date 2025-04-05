package middleware

import (
	"github.com/labstack/echo"

	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/adapter/handler"
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/usecase/authorize/token"
)

func Authorization(f handler.FunctionType) handler.FunctionType {
	return func(c echo.Context) error {
		requestToken := c.Request().Header.Get("Authorization")
		_, err := token.Validate(requestToken)
		if err != nil {
			return handler.Forbidden(c)
		}

		return f(c)
	}
}
