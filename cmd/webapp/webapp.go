package main

import (
	"log"
	"net/http"
	"os"
	"webapp/assets"
	"webapp/internal/handlers"
	"webapp/internal/models"
	"webapp/internal/utils"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Prod environment")
	}
}

func main() {
	URL := os.Getenv("URL")
	mux := http.NewServeMux()
	for path, handler := range map[string]http.Handler{
		utils.AddPre("/css/", URL): http.StripPrefix(utils.AddPre("/", URL), http.FileServer(http.FS(assets.Css))),
		utils.AddPre("/js/", URL):  http.StripPrefix(utils.AddPre("/", URL), http.FileServer(http.FS(assets.Js))),
	} {
		mux.Handle(path, utils.CacheHandler(handler))
		log.Println("Registered handler for path", path)
	}
	for path, handler := range map[string]models.HandlerFunc{
		utils.AddPre(utils.AddPre("", URL), "GET "): handlers.Index,
	} {
		mux.HandleFunc(path, utils.GenericHandler(handler, URL))
		log.Println("Registered handler for path", path)

	}
	err := http.ListenAndServe(":3000", mux)
	log.Println("Starting server on port 3000")
	if err != nil {
		log.Fatal(err)
	}

}
