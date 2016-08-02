package main

import (
	"fmt"
	"gopkg.in/couchbase/gocb.v1"
)

type Book struct {
	isbn string `json:"isbn"`
	name string `json:"name"`
	cost string `json:"cost"`
}

func main() {
	cluster, _ := gocb.Connect("couchbase://localhost")
	bucket, _ := cluster.OpenBucket("default", "")

	bucket.Upsert("u:book1",
		Book{
			isbn: "978-1-4919-1889-0",
			name: "Minecraft Modding with Forge",
			cost: "29.99",
		}, 0)

	// Get the value back
	var inBook Book
	bucket.Get("u:book1", &inBook)
	fmt.Printf("Book: %v\n", inBook)

	// Use query
	query := gocb.NewN1qlQuery("SELECT * FROM default")
	rows, err := bucket.ExecuteN1qlQuery(query, nil)
	if (err != nil) {
		fmt.Printf(err.Error())
		return;
	}
	var row interface{}
	for rows.Next(&row) {
		fmt.Printf("Row: %v", row)
	}
}

