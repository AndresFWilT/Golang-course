package handlers

import (
	"net/http"

	"github.com/AndresFWilT/afwt-clean-go-crud-test/internal/adapters/response"
	"github.com/AndresFWilT/afwt-clean-go-crud-test/internal/utils"
)

type FunctionType func(http.ResponseWriter, *http.Request)

func Forbidden(w http.ResponseWriter, r *http.Request) {
	response.GenerateError(w, utils.ValidateUUID(r.Header.Get("X-RqUID")), http.StatusUnauthorized, "Not Authorized")
	return
}
