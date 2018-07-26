package approutes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spkellydev/learngolang/src/approutes/server"
	"github.com/spkellydev/learngolang/src/middlewares"
)

// NewRouter testable router function
func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Server handled files
	r.HandleFunc("/hello", server.NonStaticResourceHandler).Methods("GET")
	r.Handle("/h", middlewares.Logger(http.HandlerFunc(server.NonStaticResourceHandler)))

	// Static files, SPA
	staticFileDir := http.Dir("./public/")                                     // Declare the static file directory to be publically show
	staticFileHandler := http.StripPrefix("/", http.FileServer(staticFileDir)) // Declare the static directory handler
	r.PathPrefix("/").Handler(staticFileHandler).Methods("GET")                // prefix path for handler

	return r
}
