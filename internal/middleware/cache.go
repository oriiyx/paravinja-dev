package middleware

import (
	"net/http"
	"strings"
	"time"
)

func CacheControl(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Cache static files
		if strings.HasPrefix(r.URL.Path, "/public/") {
			w.Header().Set("Cache-Control", "public, max-age=86400")
			w.Header().Set("Expires", time.Now().AddDate(0, 0, 1).UTC().Format(http.TimeFormat))
			w.Header().Set("Pragma", "cache")
		} else {
			// Don't cache HTML
			w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			w.Header().Set("Pragma", "no-cache")
			w.Header().Set("Expires", "0")
		}
		next.ServeHTTP(w, r)
	})
}
