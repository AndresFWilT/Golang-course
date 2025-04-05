package route

import (
	"github.com/labstack/echo"

	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/adapter/handler"
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/domain/port"
)

func Login(e *echo.Echo, storage port.PersonStorer) {
	loginHandler := handler.NewLogin(storage)

	e.POST("/api/v1/login", loginHandler.Login)
}
