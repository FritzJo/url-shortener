package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	initDatabase()
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	r.HandleFunc("/index.html", index)
	r.HandleFunc("/{shortid}", redirect)
	//r.HandleFunc("/s/{targetURL}", createShortURL)
	log.Println("Listening...")
	http.ListenAndServe(":8080", r)
}

func index(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("templates", "index.html")
	fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

        target, ok := r.URL.Query()["target"]
        base := filepath.Clean(r.Host) + "/"
        templateData := urldata{"", ""}
        if !ok || len(target[0]) < 1 {
            log.Println("Url Param 'target' is missing")
        } else {
	    short := base + string(shortenURL(target[0]))
            destination := target[0]
	    templateData = urldata{destination, short}
        }

        // Return a 404 if the template doesn't exist
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}
        layout, err := template.ParseFiles("./templates/layout.html")
	if err != nil {
		// Log the detailed error)
		log.Println(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		// Log the detailed error)
		log.Println(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if err := layout.ExecuteTemplate(w, "header", nil); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
        }

	if err := tmpl.ExecuteTemplate(w, "index", templateData); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}

        if err := layout.ExecuteTemplate(w, "footer", nil); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
        }
}

type urldata struct {
    Short string
    Target string
}

func redirect(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    shortURL := vars["shortid"]
    destinationURL := string(resolveShortURL(shortURL))
    fmt.Println("[info] destination: " + destinationURL)
    http.Redirect(w, r, destinationURL, 301)
}
