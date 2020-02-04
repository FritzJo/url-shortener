package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

type shortURL struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

func shorten(w http.ResponseWriter, r *http.Request) {
	target, _ := ioutil.ReadAll(r.Body)
	targetUrl := string(target)
	base := filepath.Clean(r.Host) + "/"
	if targetUrl == "" {
		log.Println("[INFO] Empty target data!")
		return
	}
	shortUrl := GetMD5Hash(targetUrl)
	storeURL(targetUrl, shortUrl)

	short := base + string(shortUrl)
	responsedata := shortURL{string(target), short}
	json.NewEncoder(w).Encode(responsedata)
}
