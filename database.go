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

func resolveShortURL(shortURL string) (val string) {
	log.Println("[INFO] Reading from bucket: " + shortURL)
	db, err := bolt.Open("urls.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	//fmt.Println("[info] Opened db")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("[info] reading bucket...")
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("urls"))
		//fmt.Println("[info] bucket info: " + string(b.Stats().Depth) + "!")
		val = string(b.Get([]byte(shortURL)))
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("[info] target URL: " + val)
	return val
}

func storeURL(targetURL string, shortURL string) {
	log.Println("[INFO] Writing to bucket: " + targetURL)
	db, err := bolt.Open("urls.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("urls"))
		err := b.Put([]byte(shortURL), []byte(targetURL))
		return err
	})

	if err != nil {
		log.Fatal(err)
	}
}
