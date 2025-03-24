package product

import (
	"errors"
	"time"
)

type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Models []Model

type Storager interface {
	Create(*Model) error
	Update(*Model) error
	GetAll() (Models, error)
	GetById(uint) (*Model, error)
	Delete(uint) error
}

type Service struct {
	storage Storager
}

func NewProductService(s Storager) *Service {
	return &Service{s}
}

func (s *Service) Create(m *Model) error {
	m.CreatedAt = time.Now()
	return s.storage.Create(m)
}

func (s *Service) GetAll() (Models, error) {
	return s.storage.GetAll()
}

func (s *Service) GetByID(id uint) (*Model, error) {
	return s.storage.GetById(id)
}

var (
	ErrIDNotFound = errors.New("El producto no contiene un ID")
)

func (s *Service) Update(m *Model) error {
	if m.ID == 0 {
		return ErrIDNotFound
	}
	m.UpdatedAt = time.Now()

	return s.storage.Update(m)
}

func (s *Service) Delete(id uint) error {
	return s.storage.Delete(id)
}
