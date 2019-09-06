package main

import (
	"github.com/boltdb/bolt"
	"log"
	"time"
)

func initDatabase() {
	db, err := bolt.Open("urls.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("urls"))
		if err != nil {
			log.Fatal(err)
		}
		return nil
	})

}

func resolveShortURL(shortURL string) (val []byte) {
	db, err := bolt.Open("urls.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("urls"))
		val = b.Get([]byte(shortURL))
		return nil
	})
	return val
}

func shortenURL(targetUrl string) (val []byte) {
	db, err := bolt.Open("urls.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	shortUrl := GetMD5Hash(targetUrl)
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("urls"))
		err := b.Put([]byte(shortUrl), []byte(targetUrl))
		return err
	})
	return []byte(shortUrl)
}
