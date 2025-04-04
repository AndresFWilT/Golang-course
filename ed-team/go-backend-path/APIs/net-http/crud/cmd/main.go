package main

import (
	"github.com/google/uuid"
	"net/http"
	"os"

	"github.com/AndresFWilT/afwt-clean-go-logger/console"

	"github.com/AndresFWilT/afwt-clean-go-crud-test/internal/adapters/routes"
	"github.com/AndresFWilT/afwt-clean-go-crud-test/internal/usecase/saveInMemory"
)

func main() {
	serverUUID := uuid.New()
	store := saveInMemory.NewMemory()
	mux := http.NewServeMux()

	routes.PersonRoutes(mux, &store)
	console.Log.Success(serverUUID.String(), "Starting server, running on port: %v", 9080)
	err := http.ListenAndServe(":9080", mux)
	if err != nil {
		console.Log.Error(serverUUID.String(), "Failing to start server, error: %v", err)
		os.Exit(1)
	}
}
