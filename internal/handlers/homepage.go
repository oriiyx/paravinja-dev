package handlers

import (
	"log"
	"net/http"

	"github.com/oriiyx/paravinja-dev/views/home"
)

// HomepageIndex serves homepage
func HomepageIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := home.Index().Render(r.Context(), w)
		if err != nil {
			log.Println("Error occurred in rendering homepage: ", err)
			return
		}
	}
}
