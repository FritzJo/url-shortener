package main

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "html/template"
  "path/filepath"
  "os"
  "fmt"
)

func main() {
  r := mux.NewRouter()
  r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
  r.HandleFunc("/index.html", index)
  r.HandleFunc("/{shortid}", redirect)
  log.Println("Listening...")
  http.ListenAndServe(":8080", r)
}

func index(w http.ResponseWriter, r *http.Request) {
    lp := filepath.Join("templates", "index.html")
    fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

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

    tmpl, err := template.ParseFiles(lp, fp)
    if err != nil {
    // Log the detailed error
      log.Println(err.Error())
    // Return a generic "Internal Server Error" message
      http.Error(w, http.StatusText(500), 500)
      return
    }

    if err := tmpl.ExecuteTemplate(w, "index", nil); err != nil {
      log.Println(err.Error())
      http.Error(w, http.StatusText(500), 500)
    }
}

type WebData struct {
    destURL string
}

func redirect(w http.ResponseWriter, r *http.Request) {
   // vars := mux.Vars(r)
   // wd := WebData{
    //  destURL: vars["shortid"],
    //}
  vars := mux.Vars(r)
  destinationURL := "null"
  shortURL := vars["shortid"]
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "Destination URL: %v\nShort URL: %v", destinationURL, shortURL)
}
