package handlers

import (
	"net/http"

	"github.com/AndresFWilT/afwt-clean-go-logger/console"
	"github.com/labstack/echo"

	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/adapters/response"
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/domain/models"
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/domain/ports"
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/usecase/authorize/tokens"
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/utils"
)

type Login struct {
	storage ports.PersonStorager
}

func NewLogin(storage ports.PersonStorager) Login {
	return Login{storage: storage}
}

func (l Login) Login(c echo.Context) error {
	requestUUID := utils.ValidateUUID(c.Request().Header.Get("X-RqUID"))
	console.Log.Info(requestUUID, "Entering Login Handler, Body: %v, Headers: %v", c.Request().Body, c.Request().Header)

	data := models.Login{}
	err := c.Bind(&data)
	if err != nil {
		return response.GenerateError(c, requestUUID, http.StatusBadRequest, err.Error())
	}

	// TODO: validate password and user via login use case and repository with db
	// The following logic must also be on the use case login and return the error.
	if !isLoginValid(&data) {
		return response.GenerateError(c, requestUUID, http.StatusUnauthorized, "Not Authorized")
	}

	token, err := tokens.GenerateToken(&data)
	if err != nil {
		return response.GenerateError(c, requestUUID, http.StatusInternalServerError, "Cannot create token")
	}
	dataToken := models.Token{Token: token}
	return response.Generate(c, requestUUID, http.StatusOK, "logged in and returning token", dataToken)
}

// mocking logic must perform a better approach
func isLoginValid(data *models.Login) bool {
	return data.Email == "something@gmail.com" && data.Password == "123456"
}
