package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spkellydev/learngolang/models"
)

// GetBirdHandler is the json endpoint for birds
func GetBirdHandler(w http.ResponseWriter, r *http.Request) {
	// defined in store.go
	// initialized eariler in application
	birds, err := models.BirdStore.GetBirds()
	birdListBytes, err := json.Marshal(birds)

	// if there is an error
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
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
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bird.Species = r.Form.Get("species")
	bird.Description = r.Form.Get("description")

	err = models.BirdStore.CreateBird(&bird)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
