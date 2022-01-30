package normalization

import (
	"database/sql"
	"fmt"
)

func AsyncNormalization(rows *sql.Rows, db *sql.DB) {
	rowsAmount := 0
	sync := make(chan bool)
	for rows.Next() {
		var curRow tableRow
		rows.Scan(&curRow.id, &curRow.phoneNumber, &curRow.firstName, &curRow.lastName)
		curRow.phoneNumber = normalizeNumber(curRow.phoneNumber)
		go func() {
			fmt.Println("Query runs")
			_, err := db.Query("UPDATE phones SET PhoneNumber = $1 WHERE Id = $2", curRow.phoneNumber, curRow.id)
			if err != nil {
				_, _ = db.Query("DELETE FROM phones WHERE Id = $1", curRow.id)
			}
			sync <- true
		}()
		rowsAmount++
	}
	for i := 0; i < rowsAmount; i++ {
		<-sync
	}
}
