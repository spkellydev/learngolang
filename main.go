package main

import (
	"database/sql"
	"net/http"

	"github.com/spkellydev/learngolang/approutes"
	"github.com/spkellydev/learngolang/models"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// testable router function
func newRouter() *mux.Router {
	r := approutes.NewRouter()
	return r
}

func main() {
	connString := "password=password dbname=bird_encyclopedia sslmode=disable"
	Db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}
	err = Db.Ping()
	if err != nil {
		panic(err)
	}

	models.InitStore(&models.DbStore{Db: Db})

	r := newRouter()
	http.ListenAndServe(":8080", r)
}
