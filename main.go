package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spkellydev/learngolang/approutes"
	"github.com/spkellydev/learngolang/db"

	"github.com/gorilla/mux"
	dotEnv "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// testable router function
func newRouter() *mux.Router {
	r := approutes.NewRouter()
	return r
}

// connection to psql database
func connect() {
	dbpass := os.Getenv("DBPASS") // godotenv
	dbname := os.Getenv("DBNAME") // godotenv

	// create a connection string to the database
	connString := fmt.Sprintf("password=%s dbname=%s sslmode=%s", dbpass, dbname, "disable")
	Db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	err = Db.Ping()
	if err != nil {
		panic(err)
	}

	// initialize the store
	db.InitStore(&db.Store{Db: Db})
}

func main() {
	err := dotEnv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connect()

	r := newRouter()
	http.ListenAndServe(":8000", r)
}
