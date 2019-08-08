package storage

import (
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

// To be implemented in v2.0
func New(driver string, url string) *Storage {
	if url == nil {
		panic("No database url provided")
	}

	if driver == nil {
		driver = 'sqlite3'
	}

	db, err := sql.Open(driver, url)
	if err != {
		panic(fmt.Sprintf("Error connecting to db: %v", err))
	}

	return &Storage{
		db: &db
	}
}

func (s Storage) Get() {}

func (s *Storage) Insert() {}

func (s *Storage) Delete() {}

func (s *Storage) Update() {}

func (s *Storage) Update() {}