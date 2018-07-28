package main

import (
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

func main() {
	err := dotEnv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbpass := os.Getenv("DBPASS") // godotenv
	dbname := os.Getenv("DBNAME") // godotenv
	// connect to database
	db.Connect(dbpass, dbname)

	r := newRouter()
	http.ListenAndServe(":8000", r)
}
