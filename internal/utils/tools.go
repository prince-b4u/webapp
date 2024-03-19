package utils

import (
	"net/http"
	"webapp/internal/models"
)

func GenericHandler(handler models.HandlerFunc, url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		handler(url, path, w, r)
	}
}

func AddPre(suffix, url string) string {
	return url + suffix
}

func CacheHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//w.Header().Set("Cache-Control", "public, max-age=6000 s-maxage=2592000, must-revalidate")
		h.ServeHTTP(w, r)
	})
}
