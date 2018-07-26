package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spkellydev/learngolang/src/approutes"
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
