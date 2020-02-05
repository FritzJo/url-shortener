package main

import (
	"crypto/md5"
	"encoding/hex"
)

// GetMD5Hash Generate MD5 hash of a given string
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
