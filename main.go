package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	// Connect to the DB, panic if failed

	connstring := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_DATABASE"))
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		fmt.Println(`Could not connect to db`)
		panic(err)
	}
	defer db.Close()

	t := time.Now()
	createstmt := fmt.Sprintf("CREATE TABLE TEST_TABLE_%s(counter varchar);", t.Format("20060102150405"))

	_, err = db.Exec(createstmt)
	if err != nil {
		panic(err)
	}
	insertstmt := fmt.Sprintf("INSERT INTO TEST_TABLE_%s(counter) values ('test record');", t.Format("20060102150405"))
	_, err = db.Exec(insertstmt)
	if err != nil {
		panic(err)
	}
	selectstmt := fmt.Sprintf("SELECT * FROM TEST_TABLE_%s;", t.Format("20060102150405"))
	rows, err := db.Query(selectstmt)
	if err != nil {
		panic(err)
	}

	var col1 string
	for rows.Next() {
		rows.Scan(&col1)
		fmt.Println(col1)
	}
}
