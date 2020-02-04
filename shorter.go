package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

type ShortUrl struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

func shorten(w http.ResponseWriter, r *http.Request) {
	target, _ := ioutil.ReadAll(r.Body)
	base := filepath.Clean(r.Host) + "/"
	short := base + string(shortenURL(string(target)))
	fmt.Println("t: " + string(target))
	fmt.Println("s: " + short)
	responsedata := ShortUrl{string(target), short}
	json.NewEncoder(w).Encode(responsedata)
}
