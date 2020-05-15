package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	initDatabase()
	r := mux.NewRouter()
	r.HandleFunc("/{shortid}", redirect)
	r.HandleFunc("/", index)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	r.HandleFunc("/api/v1/short", shorten).Methods("POST")

	port := os.Getenv("URL_PORT")

	if port == "" {
		log.Println("[INFO] Listening on port 8080...")
		err := http.ListenAndServe(":8080", r)
		if err != nil {
			panic(err)
		}
	} else {
		log.Println("[INFO] Listening on port " + port + "...")
		err := http.ListenAndServe(":"+port, r)
		if err != nil {
			panic(err)
		}
	}

}

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["shortid"]
	destinationURL := string(resolveShortURL(shortURL))
	http.Redirect(w, r, destinationURL, 301)
}
