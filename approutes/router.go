package approutes

import (
	"net/http"

	"github.com/spkellydev/learngolang/approutes/api"
	"github.com/spkellydev/learngolang/approutes/middlewares"

	"github.com/gorilla/mux"
)

func docs(r *mux.Router) {
	// Documentation API
	r.Handle("/api/docs", middlewares.Logger(http.HandlerFunc(api.GetDocsHandler))).Methods("GET")
	r.Handle("/api/docs", middlewares.Logger(http.HandlerFunc(api.CreateDocHandler))).Methods("POST")
	r.Handle("/api/docs/{id}", middlewares.Logger(http.HandlerFunc(api.GetDocHandler))).Methods("GET")
	r.Handle("/api/docs/{id}/update", middlewares.Logger(http.HandlerFunc(api.UpdateDocHandler))).Methods("PUT")
	r.Handle("/api/docs/{id}/destroy", middlewares.Logger(http.HandlerFunc(api.DeleteDocHandler))).Methods("DELETE")
}

func errors(r *mux.Router) {
	r.Handle("/404", middlewares.Logger(http.HandlerFunc(api.Get404))).Methods("GET")
	r.Handle("/500", middlewares.Logger(http.HandlerFunc(api.Get500))).Methods("GET")
}

// NewRouter testable router function
func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Server handled files
	// example: api
	docs(r)
	errors(r)

	// Static files, SPA
	staticFileDir := http.Dir("./public/")                                     // Declare the static file directory to be publically show
	staticFileHandler := http.StripPrefix("/", http.FileServer(staticFileDir)) // Declare the static directory handler
	r.PathPrefix("/").Handler(staticFileHandler).Methods("GET")                // prefix path for handler

	return r
}
