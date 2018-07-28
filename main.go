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

func newRouter() *mux.Router { // testable router function
	r := approutes.NewRouter() // approutes/router.go
	return r                   // release the router
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

	err = Db.Ping() // check connection
	if err != nil {
		panic(err)
	}

	// initialize the store with Db connection
	db.InitStore(&db.Store{Db: Db})
}

func main() {
	err := dotEnv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// connect to database
	connect()

	r := newRouter()
	http.ListenAndServe(":8000", r)
}
