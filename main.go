package main

import (
	"net/http"

	"approutes"

	"github.com/gorilla/mux"
)

// testable router function
func newRouter() *mux.Router {
	r := approutes.NewRouter()
	return r
}

func main() {
	r := newRouter()
	http.ListenAndServe(":8000", r)
}
