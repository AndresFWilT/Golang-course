package main

import (
	"os"

	"github.com/AndresFWilT/afwt-clean-go-logger/console"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/adapter/route"
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/usecase/authorize/certificate"
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/usecase/memory"
)

func main() {
	// TODO
	// move the server configuration to the internal/infrastructure/server layer, also put all configurations there.
	serverUUID := uuid.New()
	err := godotenv.Load()
	if err != nil {
		console.Log.Error(serverUUID.String(), "Error loading .env file: %v", err)
		os.Exit(1)
	}

	err = certificate.Load(os.Getenv("PRIVATE_CERTIFICATE_KEY_PATH"), os.Getenv("PUBLIC_CERTIFICATE_KEY_PATH"))
	if err != nil {
		console.Log.Error(serverUUID.String(), "cannot load certificate: %v", err)
		os.Exit(1)
	}
	store := memory.NewMemory()
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	route.Person(e, &store)
	route.Login(e, &store)
	console.Log.Success(serverUUID.String(), "Starting server with echo, running on port: %v", 9080)
	err = e.Start(":9080")
	if err != nil {
		console.Log.Error(serverUUID.String(), "Failing to start server with echo, error: %v", err)
		os.Exit(1)
	}
}
