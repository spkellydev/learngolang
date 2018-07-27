package main

import (
	"net/http"

	"github.com/spkellydev/learngolang/approutes"

	"github.com/gorilla/mux"
)

// testable router function
func newRouter() *mux.Router {
	r := approutes.NewRouter()
	return r
}

func main() {
	r := newRouter()
	http.ListenAndServe(":8080", r)
}
