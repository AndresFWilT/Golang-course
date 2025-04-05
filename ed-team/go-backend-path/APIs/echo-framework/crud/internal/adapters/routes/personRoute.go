package routes

import (
	"github.com/labstack/echo"
	
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/adapters/handlers"
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/adapters/middlewares"
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/domain/ports"
)

func PersonRoutes(c *echo.Echo, storage ports.PersonStorager) {
	personHandler := handlers.NewPerson(storage)

	persons := c.Group("/api/v1/persons")
	persons.POST("", echo.HandlerFunc(middlewares.Authorization(personHandler.Create)))
	persons.GET("", personHandler.GetAll)
	persons.GET("/:id", personHandler.GetById)
	persons.PUT("/:id", personHandler.Update)
	persons.DELETE("/:id", personHandler.Delete)
}
