package backend

import (
	"fmt"
	"net/http"
)

// DocsBackendUpdateHandler will handle when the client wants to update content through the backend of the website
func DocsBackendUpdateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello there")
}
