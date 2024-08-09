package handlers

import (
	"log"
	"net/http"

	about "github.com/oriiyx/paravinja-dev/views/about"
)

// AboutIndex serves contact page
func AboutIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := about.Index().Render(r.Context(), w)
		if err != nil {
			log.Println("Error occurred in rendering homepage: ", err)
			return
		}
	}
}
