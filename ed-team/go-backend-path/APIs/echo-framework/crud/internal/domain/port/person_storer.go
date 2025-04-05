package port

import "github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/domain/model"

type PersonStorer interface {
	CreatePerson(uuid string, person *model.Person) error
	UpdatePerson(uuid string, id int, person *model.Person) error
	DeletePerson(uuid string, id int) error
	GetPersonById(uuid string, id int) (model.Person, error)
	GetAllPersons(uuid string) (model.Persons, error)
}
