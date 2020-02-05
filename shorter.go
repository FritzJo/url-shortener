package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

type urlData struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

func shorten(w http.ResponseWriter, r *http.Request) {
	target, _ := ioutil.ReadAll(r.Body)
	targetURL := string(target)
	base := filepath.Clean(r.Host) + "/"
	if targetURL == "" {
		log.Println("[INFO] Empty target data!")
		return
	}
	shortURL := GetMD5Hash(targetURL)
	storeURL(targetURL, shortURL)

	short := base + string(shortURL)
	responsedata := urlData{string(target), short}
	json.NewEncoder(w).Encode(responsedata)
}
