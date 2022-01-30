package main_test

import (
	"database/sql"
	"fmt"
	"gotutorial/phone_number_normalizer/internal/normalization"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

func execQueryFromFile(db *sql.DB, filename string) error {
	queryTextBytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	_, err = db.Query(string(queryTextBytes))
	return err
}

func prepareSmallTest(db *sql.DB) {
	err := execQueryFromFile(db, "prepare_table.sql")
	normalization.LogErr(err)
	err = execQueryFromFile(db, "gen_small_test.sql")
	normalization.LogErr(err)
}

func prepareBigTest(db *sql.DB) {
	err := execQueryFromFile(db, "prepare_table.sql")
	normalization.LogErr(err)
	err = execQueryFromFile(db, "gen_big_test.sql")
	normalization.LogErr(err)
}

func generateRandomNumber(format int, len int) string {
	var nsb strings.Builder
	for i := 0; i < len; i++ {
		nsb.WriteByte(byte(rand.Intn(10)) + '0')
	}
	number := nsb.String()
	nsb.Reset()
	switch format {
	case 0:
		return number[0:1] + "(" + number[1:4] + ")" + number[4:]
	case 1:
		return number[0:1] + " (" + number[1:4] + ") " + number[4:6] + "-" + number[6:8] + "-" + number[8:]
	case 2:
		return number[0:1] + " " + number[1:4] + " " + number[4:6] + " " + number[6:8] + " " + number[8:]
	case 3:
		return number
	default:
		return ""
	}
}

func generateBigTestQueryFile(db *sql.DB) {
	const testSetSize = 5000 // number of rows in test data
	var queryText strings.Builder
	queryText.WriteString("INSERT INTO phones (PhoneNumber, SubscriberFirstName, SubscriberLastName) VALUES ")

	for i := 0; i < testSetSize; i++ {
		format := rand.Intn(4)
		number := generateRandomNumber(format, 10)
		queryText.WriteString(fmt.Sprintf("('%s','%s','%s')", number, "SubFirstName"+fmt.Sprint(i), "SubLastName"+fmt.Sprint(i)))
		if i < testSetSize-1 {
			queryText.WriteByte(',')
		} else {
			queryText.WriteByte(';')
		}
	}

	os.WriteFile("gen_big_test.sql", []byte(queryText.String()), 0644)
}

// this function is called a test just to launch it separately, it's not actually a test
func TestCreateBigTest(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=123123 dbname=test sslmode=disable")
	normalization.LogErr(err)
	defer db.Close()
	generateBigTestQueryFile(db)
}

func TestSyncNormalizationNoCacheSmallTest(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=123123 dbname=test sslmode=disable")
	normalization.LogErr(err)
	defer db.Close()
	timer := time.Now()
	prepareSmallTest(db)
	fmt.Printf("Test preparation took %dms\n", time.Since(timer).Milliseconds())
	timer = time.Now()
	rows, err := db.Query("SELECT * FROM phones;")
	normalization.LogErr(err)
	normalization.SyncNormalization(rows, db, "nocache")
	fmt.Printf("Normalization have run in %dms\n", time.Since(timer).Milliseconds())
}

func TestSyncNormalizationNoCacheBigTest(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=123123 dbname=test sslmode=disable")
	normalization.LogErr(err)
	defer db.Close()
	timer := time.Now()
	prepareBigTest(db)
	fmt.Printf("Test preparation took %dms\n", time.Since(timer).Milliseconds())
	timer = time.Now()
	rows, err := db.Query("SELECT * FROM phones;")
	normalization.LogErr(err)
	normalization.SyncNormalization(rows, db, "nocache")
	fmt.Printf("Normalization have run in %dms\n", time.Since(timer).Milliseconds())
}

func TestSyncNormalizationFullCacheSmallTest(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=123123 dbname=test sslmode=disable")
	normalization.LogErr(err)
	defer db.Close()
	timer := time.Now()
	prepareSmallTest(db)
	fmt.Printf("Test preparation took %dms\n", time.Since(timer).Milliseconds())
	timer = time.Now()
	rows, err := db.Query("SELECT * FROM phones;")
	normalization.LogErr(err)
	normalization.SyncNormalization(rows, db, "fullcache")
	fmt.Printf("Normalization have run in %dms\n", time.Since(timer).Milliseconds())
}

func TestSyncNormalizationFullCacheBigTest(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=123123 dbname=test sslmode=disable")
	normalization.LogErr(err)
	defer db.Close()
	timer := time.Now()
	prepareBigTest(db)
	fmt.Printf("Test preparation took %dms\n", time.Since(timer).Milliseconds())
	timer = time.Now()
	rows, err := db.Query("SELECT * FROM phones;")
	normalization.LogErr(err)
	normalization.SyncNormalization(rows, db, "fullcache")
	fmt.Printf("Normalization have run in %dms\n", time.Since(timer).Milliseconds())
}

func TestAsyncNormalizationSmallTest(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=123123 dbname=test sslmode=disable")
	normalization.LogErr(err)
	defer db.Close()
	timer := time.Now()
	prepareSmallTest(db)
	fmt.Printf("Test preparation took %dms\n", time.Since(timer).Milliseconds())
	timer = time.Now()
	rows, err := db.Query("SELECT * FROM phones;")
	normalization.LogErr(err)
	normalization.AsyncNormalization(rows, db)
	fmt.Printf("Normalization have run in %dms\n", time.Since(timer).Milliseconds())
}

func TestAsyncNormalizationBigTest(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=123123 dbname=test sslmode=disable")
	normalization.LogErr(err)
	defer db.Close()
	timer := time.Now()
	prepareBigTest(db)
	fmt.Printf("Test preparation took %dms\n", time.Since(timer).Milliseconds())
	timer = time.Now()
	rows, err := db.Query("SELECT * FROM phones;")
	normalization.LogErr(err)
	normalization.AsyncNormalization(rows, db)
	fmt.Printf("Normalization have run in %dms\n", time.Since(timer).Milliseconds())
}
