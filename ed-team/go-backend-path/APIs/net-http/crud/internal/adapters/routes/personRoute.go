package routes

import (
	"net/http"

	"github.com/AndresFWilT/afwt-clean-go-crud-test/internal/adapters/handlers"
	"github.com/AndresFWilT/afwt-clean-go-crud-test/internal/adapters/middlewares"
	"github.com/AndresFWilT/afwt-clean-go-crud-test/internal/domain/ports"
)

func PersonRoutes(mux *http.ServeMux, storage ports.PersonStorager) {
	personHandler := handlers.NewPerson(storage)

	mux.HandleFunc("/api/v1/persons/create", middlewares.Authorization(personHandler.Create))
	mux.HandleFunc("/api/v1/persons", personHandler.GetAll)
	mux.HandleFunc("/api/v1/person", personHandler.GetById)
	mux.HandleFunc("/api/v1/person/update", personHandler.Update)
	mux.HandleFunc("/api/v1/person/delete", personHandler.Delete)
}
