package route

import (
	"github.com/labstack/echo"

	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/adapter/handler"
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/adapter/middleware"
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/domain/port"
)

func Person(c *echo.Echo, storage port.PersonStorer) {
	personHandler := handler.NewPerson(storage)

	persons := c.Group("/api/v1/persons")
	persons.POST("", echo.HandlerFunc(middleware.Authorization(personHandler.Create)))
	persons.GET("", personHandler.GetAll)
	persons.GET("/:id", personHandler.GetById)
	persons.PUT("/:id", personHandler.Update)
	persons.DELETE("/:id", personHandler.Delete)
}
