package db

import (
	"database/sql"
	"fmt"
)

// Connect provides connection to psql database
func Connect(dbpass string, dbname string) {
	// create a connection string to the database
	connString := fmt.Sprintf("password=%s dbname=%s sslmode=%s", dbpass, dbname, "disable")
	Db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	err = Db.Ping() // check connection
	if err != nil {
		panic(err)
	}

	// initialize the store with Db connection
	InitStore(&Store{Db: Db})
}
