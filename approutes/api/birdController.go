package api

import (
	"encoding/json"
	"net/http"

	"github.com/spkellydev/learngolang/models"
)

var birds []models.Bird

// GetBirdHandler is the json endpoint for birds
func GetBirdHandler(w http.ResponseWriter, r *http.Request) {
	birdListBytes, err := json.Marshal(birds)

	// if there is an error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(birdListBytes)
}

// PostBirdHandler creates a bird
func PostBirdHandler(w http.ResponseWriter, r *http.Request) {
	bird := models.Bird{}

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bird.Species = r.Form.Get("species")
	bird.Species = r.Form.Get("description")

	birds = append(birds, bird)

	http.Redirect(w, r, "/", http.StatusFound)
}
