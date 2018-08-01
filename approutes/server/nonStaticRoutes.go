package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spkellydev/learngolang/db"
)

func handleRouteErr(err error) bool {
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		return false
	}
	return true
}

// Docs is
type Docs struct {
	Title string
	News  string
}

// GetDocsHandler Handles the GET request for documentation
func GetDocsHandler(w http.ResponseWriter, r *http.Request) {
	docs, err := db.DocsStore.GetAll()
	if ok := handleRouteErr(err); !ok {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

	docsByteList, err := json.Marshal(docs)
	if ok := handleRouteErr(err); !ok {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(docsByteList)
}

// CreateDocHandler handles the POST request to create documentation
func CreateDocHandler(w http.ResponseWriter, r *http.Request) {
	doc := db.Doc{}

	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// pull html input's name attribute
	doc.Name = r.Form.Get("name")
	doc.Package = r.Form.Get("package")
	doc.Type = r.Form.Get("type")

	// enter doc into database
	err = db.DocsStore.CreateOne(&doc)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/docs", http.StatusFound)
}
