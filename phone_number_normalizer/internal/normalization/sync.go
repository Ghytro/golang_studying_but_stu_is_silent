package normalization

import (
	"database/sql"
	"fmt"
	"strings"
)

func SyncNormalization(rows *sql.Rows, db *sql.DB, memmode string) {
	if !isSupportedMemoryMode(memmode) {
		logUnsupportedMemoryMode(memmode)
		return
	}
	var cache []tableRow
	for rows.Next() {
		var curRow tableRow
		rows.Scan(&curRow.id, &curRow.phoneNumber, &curRow.firstName, &curRow.lastName)
		curRow.phoneNumber = normalizeNumber(curRow.phoneNumber)
		if memmode == "nocache" {
			_, err := db.Query("UPDATE phones SET PhoneNumber = $1 WHERE Id = $2", curRow.phoneNumber, curRow.id)
			if err != nil {
				_, _ = db.Query("DELETE FROM phones WHERE Id = $1", curRow.id)
			}
		} else if memmode == "fullcache" {
			cache = append(cache, curRow)
		}
	}

	if memmode == "fullcache" {
		fmt.Println(cache)
		cache = filteredPhoneNumbers(cache)
		fmt.Println(cache)
		sendQueryFromCache(cache, db)
	}
}

func filteredPhoneNumbers(cache []tableRow) []tableRow { // leaves only unique phone numbers
	uniqueRows := make(map[string]tableRow)
	for _, r := range cache {
		uniqueRows[r.phoneNumber] = r
	}
	cache = make([]tableRow, 0, len(uniqueRows))
	for _, row := range uniqueRows {
		cache = append(cache, row)
	}
	return cache
}

func sendQueryFromCache(cache []tableRow, db *sql.DB) {
	const rowsLimit = 1000 // this limit is needed to divide cache into separate parts if replacing a table in single query is too much
	var queryText strings.Builder
	counter := rowsLimit
	queryText.WriteString("TRUNCATE TABLE phones; INSERT INTO phones VALUES ")
	for i, r := range cache {
		queryText.WriteString(fmt.Sprintf("(%d, '%s', '%s', '%s')", r.id, r.phoneNumber, r.firstName, r.lastName))
		counter--
		if counter != 0 && i != len(cache)-1 {
			queryText.WriteByte(',')
		} else {
			queryText.WriteByte(';')
			counter = rowsLimit

			db.Query(queryText.String())
			fmt.Println(queryText.String())
			queryText.Reset()
			queryText.WriteString("INSERT INTO phones VALUES ")
		}
	}
}
