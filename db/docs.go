package db

// Doc is a model who's responsibility it is to return the documentation table back to the application
type Doc struct {
	ID      int    // serial primary key
	Name    string // varchar 256
	Type    string // varchar 256
	Package string // varchar 256
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
	_, err := store.Db.Query("INSERT INTO documentation(name, type, package) VALUES ($1, $2, $3)", doc.Name, doc.Type, doc.Package)
	return err
}

// GetOne retreives a row from the database where the id is equivalent
func (store *Store) GetOne(id int) (*Doc, error) {
	// Query the database for an id
	row, err := store.Db.Query("SELECT * FROM documentation WHERE id IN ($1)", id)
	if err != nil {
		return nil, err
	}

	doc := &Doc{}
	for row.Next() {
		if err := row.Scan(&doc.ID, &doc.Name, &doc.Package, &doc.Type); err != nil {
			return nil, err
		}
	}

	return doc, nil
}

// DeleteOne removes a doc from the db
func (store *Store) DeleteOne(id int) error {
	_, err := store.Db.Query("DELETE FROM documentation WHERE id in ($1)", id)
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
		if err := rows.Scan(&doc.ID, &doc.Name, &doc.Package, &doc.Type); err != nil {
			return nil, err
		}

		docs = append(docs, doc)
	}

	return docs, nil
}
