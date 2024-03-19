package models

import "net/http"

type HandlerFunc func(url, path string, w http.ResponseWriter, r *http.Request)
