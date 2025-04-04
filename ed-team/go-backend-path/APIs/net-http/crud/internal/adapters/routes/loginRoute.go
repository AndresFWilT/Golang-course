package routes

import (
	"net/http"

	"github.com/AndresFWilT/afwt-clean-go-crud-test/internal/adapters/handlers"
	"github.com/AndresFWilT/afwt-clean-go-crud-test/internal/domain/ports"
)

func LoginRoutes(mux *http.ServeMux, storage ports.PersonStorager) {
	loginHandler := handlers.NewLogin(storage)

	mux.HandleFunc("/api/v1/login", loginHandler.Login)
}
