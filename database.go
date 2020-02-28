package main

import (
	"log"
	"time"

	"github.com/boltdb/bolt"
)

func initDatabase() {
	db, err := bolt.Open("urls.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("urls"))
		if err != nil {
			log.Fatal(err)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

}

func accessBucket(val1 string, val2 string, mode string) string {
	db, err := bolt.Open("urls.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	result := ""
	//fmt.Println("[info] reading bucket...")
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("urls"))
		if mode == "put" {
			log.Println("[INFO] Writing to bucket: " + val2)
			err := b.Put([]byte(val1), []byte(val2))

			if err != nil {
				log.Fatal(err)
			}
		} else if mode == "get" {
			log.Println("[INFO] Reading from bucket: " + val1)
			val := string(b.Get([]byte(val1)))
			result = val
		}
		return nil
	})
	return result

}

func resolveShortURL(shortURL string) (val string) {
	return accessBucket(shortURL, "", "get")
}

func storeURL(targetURL string, shortURL string) {
	accessBucket(targetURL, shortURL, "put")
}
