package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	initDatabase()
	r := mux.NewRouter()
	r.HandleFunc("/{shortid}", redirect)
	r.HandleFunc("/", index)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	r.HandleFunc("/api/v1/short", shorten).Methods("POST")

	log.Println("[INFO] Listening...")
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		panic(err)
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
