package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AndresFWilT/afwt-clean-go-logger/console"

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

func (l Login) Login(w http.ResponseWriter, r *http.Request) {
	requestUUID := utils.ValidateUUID(r.Header.Get("X-RqUID"))
	console.Log.Info(requestUUID, "Entering Login Handler, Body: %v, Headers: %v", r.Body, r.Header)
	if r.Method != http.MethodPost {
		response.GenerateError(w, requestUUID, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	data := models.Login{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response.GenerateError(w, requestUUID, http.StatusBadRequest, err.Error())
		return
	}

	// TODO: validate password and user via login use case and repository with db
	// The following logic must also be on the use case login and return the error.
	if !isLoginValid(&data) {
		response.GenerateError(w, requestUUID, http.StatusUnauthorized, "Not Authorized")
		return
	}

	token, err := tokens.GenerateToken(&data)
	if err != nil {
		response.GenerateError(w, requestUUID, http.StatusInternalServerError, "Cannot create token")
		return
	}
	dataToken := models.Token{Token: token}
	response.Generate(w, requestUUID, http.StatusOK, "logged in and returning token", dataToken)
}

// mocking logic must perform a better approach
func isLoginValid(data *models.Login) bool {
	return data.Email == "something@gmail.com" && data.Password == "123456"
}
