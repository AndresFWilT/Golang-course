package storage

import (
	"database/sql"
	"fmt"
	"github.com/AndresFWilT/postgres-db/package/product"
	_ "github.com/lib/pq"
	"log"
	"sync"
	"time"
)

var (
	db   *sql.DB
	once sync.Once
)

// NewPostgresDb Singleton pattern db postgres connection
func NewPostgresDb() {
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", "postgres://pogotest:go1234@localhost:7530/postgresgo?sslmode=disable")
		if err != nil {
			log.Fatalf("Can't connect to postgres database: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Can't ping to postgres database: %v", err)
		}
		fmt.Println("Successfully connected to postgres database")
	})
}

// Pool unique instance for singleton
func Pool() *sql.DB {
	return db
}

func stringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}
	if null.String != "" {
		null.Valid = true
	}
	return null
}

type scanner interface {
	Scan(dest ...interface{}) error
}

func scanRowProduct(s scanner) (*product.Model, error) {
	m := &product.Model{}
	observationNull := sql.NullString{}
	updatedAtNull := sql.NullTime{}

	err := s.Scan(
		&m.ID,
		&m.Name,
		&observationNull,
		&m.Price,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return &product.Model{}, err
	}

	m.Observations = observationNull.String
	m.UpdatedAt = updatedAtNull.Time

	return m, nil
}

func timeToNull(t time.Time) sql.NullTime {
	null := sql.NullTime{Time: t}
	if !null.Time.IsZero() {
		null.Valid = true
	}
	return null
}
