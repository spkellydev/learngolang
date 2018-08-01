package db

import "database/sql"

// StoreInterface provides and interface for every model to give them default functionality
// uses Doc from docs.go
type StoreInterface interface {
	GetAll() ([]*Doc, error)
	GetOne(id int) (*Doc, error)
	CreateOne(model *Doc) error
	UpdateOne(model *Doc) error
	DeleteOne(id int) error
}

// Store holds the native go sql methods
type Store struct {
	Db *sql.DB
}

// DocsStore is the provider for documentationd data
var DocsStore StoreInterface

// InitStore is the db runner
func InitStore(s StoreInterface) {
	DocsStore = s
}
