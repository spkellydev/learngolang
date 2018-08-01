package db

// Doc is a model who's responsibility it is to return the documentation table back to the application
type Doc struct {
	ID      int
	Name    string
	Type    string
	Methods []*method
	Package string
}

type method struct {
	PackageID   int
	Name        string
	Function    string
	Description string
	ReturnType  string
	Params      []string
}

// CreateOne allows for creating one documentation row
// all types required
func (store *Store) CreateOne(doc *Doc) error {
	// Query the database for values from documentation table
	// Todo -- Automate builder
	// Todo -- doc.Methods needs to be Marshaled into string
	_, err := store.Db.Query("INSERT INTO documentation(id, name, type, methods, package) VALUES ($1, $2, $3, $4, $5)", doc.ID, doc.Name, doc.Type, doc.Methods, doc.Package)
	return err
}

// GetAll retrieves all docs from database without reqards to where any values are true
func (store *Store) GetAll() ([]*Doc, error) {
	// Query the database for all docs
	rows, err := store.Db.Query("SELECT * from documentation")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	docs := []*Doc{}
	for rows.Next() {
		// create a pointer to the current Doc
		doc := &Doc{}
		if err := rows.Scan(&doc.ID, &doc.Name, &doc.Methods, &doc.Package, &doc.Type); err != nil {
			return nil, err
		}

		docs = append(docs, doc)
	}

	return docs, nil
}
