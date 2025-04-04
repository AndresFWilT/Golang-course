package ports

import "github.com/AndresFWilT/afwt-clean-go-crud-test/internal/domain/models"

type PersonStorager interface {
	CreatePerson(uuid string, person *models.Person) error
	UpdatePerson(uuid string, id int, person *models.Person) error
	DeletePerson(uuid string, id int) error
	GetPersonById(uuid string, id int) (models.Person, error)
	GetAllPersons(uuid string) (models.Persons, error)
}
