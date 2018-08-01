package db

import "database/sql"

// StoreInterface provides and interface for every model to give them default functionality
// uses Doc from docs.go
type StoreInterface []interface {
	GetAll() ([]*Doc, error)
	// GetOne() (*Doc, error)
	CreateOne(model *Doc) error
	// UpdateOne(model *Doc) error
	// DeleteOne(model *Doc) error
}

// Store holds the native go sql methods
type Store struct {
	Db *sql.DB
}

// DocsStore is the provider for documentationd data
var DocsStore StoreInterface

// InitStore is the db runner
func InitStore(name string, s StoreInterface) {
	switch name {
	case "docs":
		DocsStore = s
	}
}
