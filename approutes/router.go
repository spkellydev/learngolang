package approutes

import (
	"net/http"

	"github.com/spkellydev/learngolang/approutes/middlewares"
	"github.com/spkellydev/learngolang/approutes/server"

	"github.com/gorilla/mux"
)

func docs(r *mux.Router) {
	// Documentation API
	r.Handle("/api/docs", middlewares.Logger(http.HandlerFunc(server.GetDocsHandler))).Methods("GET")
	r.Handle("/api/docs", middlewares.Logger(http.HandlerFunc(server.CreateDocHandler))).Methods("POST")
	r.Handle("/api/docs/{id}", middlewares.Logger(http.HandlerFunc(server.GetDocHandler))).Methods("GET")
	r.Handle("/api/docs/{id}/destroy", middlewares.Logger(http.HandlerFunc(server.DeleteDocHandler))).Methods("DELETE")
}

// NewRouter testable router function
func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Server handled files
	// example: api
	docs(r)

	// Static files, SPA
	staticFileDir := http.Dir("./public/")                                     // Declare the static file directory to be publically show
	staticFileHandler := http.StripPrefix("/", http.FileServer(staticFileDir)) // Declare the static directory handler
	r.PathPrefix("/").Handler(staticFileHandler).Methods("GET")                // prefix path for handler

	return r
}
