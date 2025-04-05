package handler

import (
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/shared/validation"
	"net/http"

	"github.com/AndresFWilT/afwt-clean-go-logger/console"
	"github.com/labstack/echo"

	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/adapter/response"
)

type FunctionType func(c echo.Context) error

func Forbidden(c echo.Context) error {
	requestUUID := validation.ValidateUUID(c.Request().Header.Get("X-RqUID"))
	console.Log.Info(requestUUID, "Returning a Forbidden response, Body: %v, Headers: %v", c.Request().Body, c.Request().Header)
	return response.GenerateError(c, requestUUID, http.StatusUnauthorized, "Not Authorized")
}
