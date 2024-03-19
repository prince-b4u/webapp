package handlers

import (
	"net/http"
	"webapp/internal/views"
)

func Index(URL, path string, w http.ResponseWriter, r *http.Request) {
	views.Index("World", r).Render(r.Context(), w)
}
