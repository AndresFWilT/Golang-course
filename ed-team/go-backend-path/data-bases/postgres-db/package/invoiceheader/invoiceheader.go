package invoiceheader

import (
	"database/sql"
	"time"
)

// Model of invoiceheader
type Model struct {
	ID        uint
	Client    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Storage interface that must implement a db storage
type Storage interface {
	CreateTx(*sql.Tx, *Model) error
}

// Service of invoiceheader
type Service struct {
	storage Storage
}

// NewService return a pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}
