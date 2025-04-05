package routes

import (
	"github.com/labstack/echo"

	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/adapters/handlers"
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/domain/ports"
)

func LoginRoutes(e *echo.Echo, storage ports.PersonStorager) {
	loginHandler := handlers.NewLogin(storage)

	e.POST("/api/v1/login", loginHandler.Login)
}
