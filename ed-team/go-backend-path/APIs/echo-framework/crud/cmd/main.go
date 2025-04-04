package main

import (
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"net/http"
	"os"

	"github.com/AndresFWilT/afwt-clean-go-logger/console"

	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/adapters/routes"
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/usecase/authorize/certificates"
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/usecase/saveInMemory"
)

func main() {
	serverUUID := uuid.New()
	err := godotenv.Load()
	if err != nil {
		console.Log.Error(serverUUID.String(), "Error loading .env file: %v", err)
		os.Exit(1)
	}

	err = certificates.LoadCertificates(os.Getenv("PRIVATE_CERTIFICATE_KEY_PATH"), os.Getenv("PUBLIC_CERTIFICATE_KEY_PATH"))
	if err != nil {
		console.Log.Error(serverUUID.String(), "cannot load certificates: %v", err)
		os.Exit(1)
	}
	store := saveInMemory.NewMemory()
	mux := http.NewServeMux()

	routes.PersonRoutes(mux, &store)
	routes.LoginRoutes(mux, &store)
	console.Log.Success(serverUUID.String(), "Starting server, running on port: %v", 9080)
	err = http.ListenAndServe(":9080", mux)
	if err != nil {
		console.Log.Error(serverUUID.String(), "Failing to start server, error: %v", err)
		os.Exit(1)
	}
}
