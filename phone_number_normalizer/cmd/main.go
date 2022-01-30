package main

import (
	"database/sql"
	"flag"
	"fmt"

	"gotutorial/phone_number_normalizer/internal/normalization"

	_ "github.com/lib/pq"
)

func main() {
	async := flag.Bool("async", false, "Specifies the normalization algorithm: should it be sync or async")
	storeCache := flag.Bool("storecache", false, "Specifies the normalization algorithm: should it store full table in cache or no")
	flag.Parse()
	db, err := sql.Open("postgres", "host=localhost user=postgres password=123123 dbname=test sslmode=disable")
	normalization.LogErr(err)
	rows, err := db.Query("SELECT * FROM phones;")
	normalization.LogErr(err)
	var memPolicy string
	if *storeCache {
		memPolicy = "fullcache"
	} else {
		memPolicy = "nocache"
	}
	if *async {
		fmt.Println("here1")
		normalization.AsyncNormalization(rows, db)
	} else {
		fmt.Println("here")
		normalization.SyncNormalization(rows, db, memPolicy)
	}
}
