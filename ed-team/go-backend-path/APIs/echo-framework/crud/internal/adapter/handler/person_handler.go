package handler

import (
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/shared/validation"
	"net/http"
	"strconv"

	"github.com/AndresFWilT/afwt-clean-go-logger/console"
	"github.com/labstack/echo"

	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/adapter/response"
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/domain/model"
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/domain/port"
)

type Person struct {
	storage port.PersonStorer
}

func NewPerson(storage port.PersonStorer) Person {
	return Person{storage: storage}
}

func (p *Person) Create(c echo.Context) error {
	requestUUID := validation.ValidateUUID(c.Request().Header.Get("X-RqUID"))
	console.Log.Info(requestUUID, "Entering Create Person Handler, Body: %v, Headers: %v", c.Request().Body, c.Request().Header)

	data := model.Person{}
	err := c.Bind(&data)

	if err != nil {
		return response.GenerateError(c, requestUUID, http.StatusBadRequest, err.Error())
	}

	err = p.storage.CreatePerson(requestUUID, &data)
	if err != nil {
		return response.GenerateError(c, requestUUID, http.StatusInternalServerError, err.Error())
	}

	return response.Generate(c, requestUUID, http.StatusCreated, "Person created", nil)
}

func (p *Person) GetAll(c echo.Context) error {
	requestUUID := validation.ValidateUUID(c.Request().Header.Get("X-RqUID"))
	console.Log.Info(requestUUID, "Entering Get All Persons Handler, Headers: %v", c.Request().Header)

	persons, err := p.storage.GetAllPersons(requestUUID)
	if err != nil {
		return response.GenerateError(c, requestUUID, http.StatusInternalServerError, err.Error())
	}

	return response.Generate(c, requestUUID, http.StatusOK, "All persons obtained", persons)
}

func (p *Person) GetById(c echo.Context) error {
	requestUUID := validation.ValidateUUID(c.Request().Header.Get("X-RqUID"))
	console.Log.Info(requestUUID, "Entering Getting Person By Id Handler, RequestParam: %v, Body: %v Headers: %v", c.QueryParams(), c.Request().Body, c.Request().Header)

	personId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.GenerateError(c, requestUUID, http.StatusBadRequest, err.Error())
	}

	person, err := p.storage.GetPersonById(requestUUID, personId)
	if err != nil {
		return response.GenerateError(c, requestUUID, http.StatusInternalServerError, err.Error())
	}

	return response.Generate(c, requestUUID, http.StatusOK, "Person obtained", person)
}

func (p *Person) Update(c echo.Context) error {
	requestUUID := validation.ValidateUUID(c.Request().Header.Get("X-RqUID"))
	console.Log.Info(requestUUID, "Entering Update Person Handler, RequestParam: %v, Body: %v Headers: %v", c.QueryParams(), c.Request().Body, c.Request().Header)

	personId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.GenerateError(c, requestUUID, http.StatusBadRequest, err.Error())
	}

	data := model.Person{}
	err = c.Bind(&data)
	if err != nil {
		return response.GenerateError(c, requestUUID, http.StatusBadRequest, err.Error())
	}

	err = p.storage.UpdatePerson(requestUUID, personId, &data)
	if err != nil {
		return response.GenerateError(c, requestUUID, http.StatusInternalServerError, err.Error())
	}

	return response.Generate(c, requestUUID, http.StatusOK, "Update person", nil)
}

func (p *Person) Delete(c echo.Context) error {
	requestUUID := validation.ValidateUUID(c.Request().Header.Get("X-RqUID"))
	console.Log.Info(requestUUID, "Entering Deleting Person By Id Handler, RequestParam: %v, Body: %v Headers: %v", c.QueryParams(), c.Request().Body, c.Request().Header)

	personId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.GenerateError(c, requestUUID, http.StatusBadRequest, err.Error())
	}

	err = p.storage.DeletePerson(requestUUID, personId)
	if err != nil {
		return response.GenerateError(c, requestUUID, http.StatusInternalServerError, err.Error())
	}

	return response.Generate(c, requestUUID, http.StatusAccepted, "Person deleted", nil)
}
