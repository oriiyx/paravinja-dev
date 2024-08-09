package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/oriiyx/paravinja-dev/views/home"
)

// PostsIndex serves homepage
func PostsIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug, err := strconv.Atoi(chi.URLParam(r, "slug"))
		if err != nil {
			return r.Write([]byte("Invalid slug"), http.StatusBadRequest)
		}

		err := home.Index().Render(r.Context(), w)
		if err != nil {
			log.Println("Error occurred in rendering homepage: ", err)
			return
		}
	}
}
