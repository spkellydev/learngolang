package server

import (
	"fmt"
	"net/http"
)

// NonStaticResourceHandler should handle server resources
// TODO - make route specific handlers
func NonStaticResourceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go docs for noobs")
}
