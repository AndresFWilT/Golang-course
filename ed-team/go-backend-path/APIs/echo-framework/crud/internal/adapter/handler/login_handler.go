package handler

import (
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/shared/validation"
	"net/http"

	"github.com/AndresFWilT/afwt-clean-go-logger/console"
	"github.com/labstack/echo"

	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/adapter/response"
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/domain/model"
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/domain/port"
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/usecase/authorize/token"
)

type Login struct {
	storage port.PersonStorer
}

func NewLogin(storage port.PersonStorer) Login {
	return Login{storage: storage}
}

func (l Login) Login(c echo.Context) error {
	requestUUID := validation.ValidateUUID(c.Request().Header.Get("X-RqUID"))
	console.Log.Info(requestUUID, "Entering Login Handler, Body: %v, Headers: %v", c.Request().Body, c.Request().Header)

	data := model.Login{}
	err := c.Bind(&data)
	if err != nil {
		return response.GenerateError(c, requestUUID, http.StatusBadRequest, err.Error())
	}

	// TODO: validate password and user via login use case and repository with db
	// The following logic must also be on the use case login and return the error.
	if !isLoginValid(&data) {
		return response.GenerateError(c, requestUUID, http.StatusUnauthorized, "Not Authorized")
	}

	retrievedToken, err := token.Generate(&data)
	if err != nil {
		return response.GenerateError(c, requestUUID, http.StatusInternalServerError, "Cannot create token")
	}
	dataToken := model.Token{Token: retrievedToken}
	return response.Generate(c, requestUUID, http.StatusOK, "logged in and returning token", dataToken)
}

// mocking logic must perform a better approach
func isLoginValid(data *model.Login) bool {
	return data.Email == "something@gmail.com" && data.Password == "123456"
}
