package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// testable router function
func newRouter() *mux.Router {
	r := mux.NewRouter()
	// Declare the static file directory to be publically show
	staticFileDir := http.Dir("./public/")
	// Declare the static directory handler
	staticFileHandler := http.StripPrefix("/", http.FileServer(staticFileDir))

	r.PathPrefix("/").Handler(staticFileHandler).Methods("GET")
	return r
}

func main() {
	r := newRouter()
	http.ListenAndServe(":80", r)
}

// useless func
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}
