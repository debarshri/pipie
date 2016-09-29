package main

import (
	"os"
	"github.com/boltdb/bolt"
	"fmt"
)

func main() {

	dbpath := os.Args[1]

	fmt.Println("The database is ",dbpath)

	db, err := bolt.Open(dbpath, 0600, nil)
	if err != nil {
		fmt.Println(err)
	}

	db.Update(func(tx *bolt.Tx) error {

		c := tx.Bucket([]byte("data"))

		if c != nil {

			cursor := c.Cursor()
			var count int = 0

			for k, _ := cursor.First(); k != nil; k, _ = cursor.Next() {

				count = count+1
			}

			fmt.Println("Total messages in the queue ",count)
		} else {
			fmt.Println("Bucket is empty")
		}

		return nil
	})
}
