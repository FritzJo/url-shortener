package main

import (
	"encoding/json"
	"hash/adler32"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
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
	shortURL := shortFunction(target)
	url, err := resolveShortURL(shortURL)
	if err != nil {
		panic(err)
	}
	for url != "" && url != targetURL {
		shortURL = shortURL + "a"
		url, err = resolveShortURL(shortURL)
		if err != nil {
			panic(err)
		}

	}
	storeURL(targetURL, shortURL)

	short := base + string(shortURL)
	responsedata := urlData{string(target), short}
	json.NewEncoder(w).Encode(responsedata)
}

func shortFunction(url []byte) string {
	adler32Int := adler32.Checksum(url)
	return strconv.FormatUint(uint64(adler32Int), 16)
}
