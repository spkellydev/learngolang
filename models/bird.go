package models

// Bird is a fucking bird
type Bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

// CreateBird will accept only a Bird struct and add it to the database
func (store *DbStore) CreateBird(bird *Bird) error {
	// don't care about the return value unless its an error
	_, err := store.Db.Query("INSERT INTO birds(species, description) VALUES ($1, $2)", bird.Species, bird.Description)
	return err
}

// GetBirds returns all birds
func (store *DbStore) GetBirds() ([]*Bird, error) {
	// Query the database for all birds and return the result
	// row object
	rows, err := store.Db.Query("SELECT species, description from birds")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	birds := []*Bird{}
	for rows.Next() {
		// create a pointer to a Bird for each row
		bird := &Bird{}
		if err := rows.Scan(&bird.Species, &bird.Description); err != nil {
			return nil, err
		}

		// append the return to the return array and repeat for next row
		birds = append(birds, bird)
	}

	return birds, nil
}
