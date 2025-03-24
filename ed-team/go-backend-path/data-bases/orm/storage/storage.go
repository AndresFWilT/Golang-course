package storage

import (
	"fmt"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db   *gorm.DB
	once sync.Once
)

// Driver of storage
type Driver string

// Drivers
const (
	Postgres Driver = "POSTGRES"
)

// New create the connection with db
func New(d Driver) {
	switch d {
	case Postgres:
		newPostgresDB()
	}
}

func newPostgresDB() {
	once.Do(func() {
		var err error
		db, err = gorm.Open("postgres", "postgres://pogotest:go1234@localhost:7530/postgresgo?sslmode=disable")
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		fmt.Println("Connected to postgres via GORM")
	})
}

// DB return a unique instance of db
func DB() *gorm.DB {
	return db
}
