package models

import "database/sql"

// Store represents an interface with two methods,
// add a bird or get all existing birds
// each method returns an error in case something goes wrong
type Store interface {
	CreateBird(bird *Bird) error
	GetBirds() ([]*Bird, error)
}

// DbStore holds the db
type DbStore struct {
	Db *sql.DB
}

// BirdStore Holds the store
var BirdStore Store

// InitStore Initializes our store on the package level
func InitStore(s Store) {
	BirdStore = s
}
