package handlers

import (
	"log"
	"net/http"

	contact "github.com/oriiyx/paravinja-dev/views/contact"
)

// ContactIndex serves contact page
func ContactIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := contact.Index().Render(r.Context(), w)
		if err != nil {
			log.Println("Error occurred in rendering homepage: ", err)
			return
		}
	}
}
