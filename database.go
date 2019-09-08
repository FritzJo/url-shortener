package main

import (
	"fmt"
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

func resolveShortURL(shortURL string) (val string) {
	db, err := bolt.Open("urls.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	fmt.Println("[info] Opened db")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[info] reading bucket...")
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("urls"))
		fmt.Println("[info] bucket info: " + string(b.Stats().Depth) + "!")
		val = string(b.Get([]byte(shortURL)))
		return nil
	})
	fmt.Println("[info] target URL: " + val)
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
