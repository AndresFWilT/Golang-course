package handlers

import (
	"net/http"

	"github.com/AndresFWilT/afwt-clean-go-logger/console"
	"github.com/labstack/echo"

	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/adapters/response"
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/utils"
)

type FunctionType func(c echo.Context) error

func Forbidden(c echo.Context) error {
	requestUUID := utils.ValidateUUID(c.Request().Header.Get("X-RqUID"))
	console.Log.Info(requestUUID, "Returning a Forbidden response, Body: %v, Headers: %v", c.Request().Body, c.Request().Header)
	return response.GenerateError(c, requestUUID, http.StatusUnauthorized, "Not Authorized")
}
