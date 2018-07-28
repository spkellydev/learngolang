package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/spkellydev/learngolang/db"
)

// handleRouteErr accepts the ResponseWrite and the status code to abstract
// some of the bad parts of Go error handling
func handleRouteErr(err error, w http.ResponseWriter, r http.Request, statusCode int) {
	ok := func() bool {
		if err != nil {
			return false // trigger panic
		}
		return true // continue onward
	}

	if !ok() {
		ErrorRequestHandler(w, &r, statusCode)
	}

	return
}

// Docs is
type Docs struct {
	Title string
	News  string
}

// CreateDocHandler handles the POST request to create documentation
func CreateDocHandler(w http.ResponseWriter, r *http.Request) {
	doc := db.Doc{}

	err := r.ParseForm()
	handleRouteErr(err, w, *r, http.StatusBadRequest)

	// pull html input's name attribute
	doc.Name = r.Form.Get("name")
	doc.Package = r.Form.Get("package")
	doc.Type = r.Form.Get("type")

	// enter doc into database
	err = db.DocsStore.CreateOne(&doc)
	handleRouteErr(err, w, *r, http.StatusInternalServerError)

	http.Redirect(w, r, "/docs", http.StatusFound)
}

// GetDocHandler Handles the GET request for documentation
func GetDocHandler(w http.ResponseWriter, r *http.Request) {
	query := mux.Vars(r)["id"]     // get id from route parameter
	id, err := strconv.Atoi(query) // convert route parameter into int from string
	handleRouteErr(err, w, *r, http.StatusInternalServerError)

	docs, err := db.DocsStore.GetOne(id)
	handleRouteErr(err, w, *r, http.StatusInternalServerError)

	docsByteList, err := json.Marshal(docs)
	handleRouteErr(err, w, *r, http.StatusInternalServerError)

	w.Header().Set("Content-Type", "application/json")
	w.Write(docsByteList)
}

// GetDocsHandler Handles the GET request for documentation
func GetDocsHandler(w http.ResponseWriter, r *http.Request) {
	docs, err := db.DocsStore.GetAll()
	handleRouteErr(err, w, *r, http.StatusInternalServerError)

	docsByteList, err := json.Marshal(docs)
	handleRouteErr(err, w, *r, http.StatusInternalServerError)

	w.Header().Set("Content-Type", "application/json")
	w.Write(docsByteList)
}

// UpdateDocHandler will update
func UpdateDocHandler(w http.ResponseWriter, r *http.Request) {
	doc := &db.Doc{}
	if err := r.ParseForm(); err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

	// convert query request to appropriate data type
	// int
	fmt.Printf("EE: %s", r.Form.Get("type"))

	query := r.Form.Get("id")
	id, err := strconv.Atoi(query)
	if err != nil {
		fmt.Printf("Error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO -- marshal request into struct
	doc.ID = id
	doc.Type = r.Form.Get("type")
	doc.Package = r.Form.Get("package")
	doc.Name = r.Form.Get("name")

	err = db.DocsStore.UpdateOne(doc)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/api/docs", http.StatusAccepted)
}

// DeleteDocHandler will delete one item from the documentation database
func DeleteDocHandler(w http.ResponseWriter, r *http.Request) {
	query := mux.Vars(r)["id"]
	id, err := strconv.Atoi(query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

	err = db.DocsStore.DeleteOne(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

	http.Redirect(w, r, "/api/docs", http.StatusAccepted)
}
