package saveInMemory

import (
	"fmt"

	"github.com/AndresFWilT/afwt-clean-go-logger/console"

	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/domain/models"
)

type Memory struct {
	CurrentId int
	Persons   map[int]models.Person
}

func NewMemory() Memory {
	return Memory{
		CurrentId: 0,
		Persons:   make(map[int]models.Person),
	}
}

func (m *Memory) CreatePerson(uuid string, person *models.Person) error {
	if person == nil {
		return models.ErrPersonCannotBeNil
	}

	console.Log.Info(uuid, "Creating person: %v", person)
	m.CurrentId++
	m.Persons[m.CurrentId] = *person
	return nil
}

func (m *Memory) UpdatePerson(uuid string, id int, person *models.Person) error {
	if person == nil {
		return models.ErrPersonCannotBeNil
	}

	if _, ok := m.Persons[id]; !ok {
		return fmt.Errorf("person with id %d: %w", id, models.ErrPersonIdNotFound)
	}

	console.Log.Info(uuid, "Updating person: %v", person)
	m.Persons[id] = *person

	return nil
}

func (m *Memory) DeletePerson(uuid string, id int) error {
	if _, ok := m.Persons[id]; !ok {
		return fmt.Errorf("person with id %d: %w", id, models.ErrPersonIdNotFound)
	}

	console.Log.Info(uuid, "Deleting person with id: %v", id)
	delete(m.Persons, id)

	return nil
}

func (m *Memory) GetPersonById(uuid string, id int) (models.Person, error) {
	console.Log.Info(uuid, "Getting Person By id: %v", id)

	foundPerson, ok := m.Persons[id]
	if !ok {
		return foundPerson, fmt.Errorf("person with id %d: %w", id, models.ErrPersonIdNotFound)
	}

	console.Log.Info(uuid, "Person got successfully")
	return foundPerson, nil
}

func (m *Memory) GetAllPersons(uuid string) (models.Persons, error) {
	console.Log.Info(uuid, "Getting all persons")

	var allPersons models.Persons
	for _, person := range m.Persons {
		allPersons = append(allPersons, person)
	}
	return allPersons, nil
}
