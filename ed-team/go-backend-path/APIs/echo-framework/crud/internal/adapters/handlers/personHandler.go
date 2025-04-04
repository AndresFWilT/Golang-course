package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AndresFWilT/afwt-clean-go-logger/console"

	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/adapters/response"
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/domain/models"
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/domain/ports"
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/utils"
)

type Person struct {
	storage ports.PersonStorager
}

func NewPerson(storage ports.PersonStorager) Person {
	return Person{storage: storage}
}

func (p *Person) Create(w http.ResponseWriter, r *http.Request) {
	requestUUID := utils.ValidateUUID(r.Header.Get("X-RqUID"))
	console.Log.Info(requestUUID, "Entering Create Person Handler, Body: %v, Headers: %v", r.Body, r.Header)
	if r.Method != http.MethodPost {
		response.GenerateError(w, requestUUID, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	data := models.Person{}
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		response.GenerateError(w, requestUUID, http.StatusBadRequest, err.Error())
		return
	}

	err = p.storage.CreatePerson(requestUUID, &data)
	if err != nil {
		response.GenerateError(w, requestUUID, http.StatusInternalServerError, err.Error())
		return
	}

	response.Generate(w, requestUUID, http.StatusCreated, "Person created", nil)

}

func (p *Person) GetAll(w http.ResponseWriter, r *http.Request) {
	requestUUID := utils.ValidateUUID(r.Header.Get("X-RqUID"))
	console.Log.Info(requestUUID, "Entering Get All Persons Handler, Headers: %v", r.Header)
	if r.Method != http.MethodGet {
		response.GenerateError(w, requestUUID, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	persons, err := p.storage.GetAllPersons(requestUUID)
	if err != nil {
		response.GenerateError(w, requestUUID, http.StatusInternalServerError, err.Error())
		return
	}

	response.Generate(w, requestUUID, http.StatusOK, "All persons obtained", persons)
}

func (p *Person) GetById(w http.ResponseWriter, r *http.Request) {
	requestUUID := utils.ValidateUUID(r.Header.Get("X-RqUID"))
	console.Log.Info(requestUUID, "Entering Getting Person By Id Handler, RequestParam: %v, Body: %v Headers: %v", r.URL.Query().Encode(), r.Body, r.Header)
	if r.Method != http.MethodGet {
		response.GenerateError(w, requestUUID, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	personId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response.GenerateError(w, requestUUID, http.StatusBadRequest, err.Error())
		return
	}

	person, err := p.storage.GetPersonById(requestUUID, personId)
	if err != nil {
		response.GenerateError(w, requestUUID, http.StatusInternalServerError, err.Error())
		return
	}

	response.Generate(w, requestUUID, http.StatusOK, "Person obtained", person)
}

func (p *Person) Update(w http.ResponseWriter, r *http.Request) {
	requestUUID := utils.ValidateUUID(r.Header.Get("X-RqUID"))
	console.Log.Info(requestUUID, "Entering Update Person Handler, RequestParam: %v, Body: %v Headers: %v", r.URL.Query().Encode(), r.Body, r.Header)
	if r.Method != http.MethodPut {
		response.GenerateError(w, requestUUID, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	personId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response.GenerateError(w, requestUUID, http.StatusBadRequest, err.Error())
		return
	}

	data := models.Person{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response.GenerateError(w, requestUUID, http.StatusBadRequest, err.Error())
		return
	}

	err = p.storage.UpdatePerson(requestUUID, personId, &data)
	if err != nil {
		response.GenerateError(w, requestUUID, http.StatusInternalServerError, err.Error())
		return
	}

	response.Generate(w, requestUUID, http.StatusOK, "Update person", nil)
}

func (p *Person) Delete(w http.ResponseWriter, r *http.Request) {
	requestUUID := utils.ValidateUUID(r.Header.Get("X-RqUID"))
	console.Log.Info(requestUUID, "Entering Deleting Person By Id Handler, RequestParam: %v, Body: %v Headers: %v", r.URL.Query().Encode(), r.Body, r.Header)
	if r.Method != http.MethodDelete {
		response.GenerateError(w, requestUUID, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	personId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response.GenerateError(w, requestUUID, http.StatusBadRequest, err.Error())
		return
	}

	err = p.storage.DeletePerson(requestUUID, personId)
	if err != nil {
		response.GenerateError(w, requestUUID, http.StatusInternalServerError, err.Error())
		return
	}

	response.Generate(w, requestUUID, http.StatusAccepted, "Person deleted", nil)
}
