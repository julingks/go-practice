package main

import (
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbName := flag.String("db", "testdb", "database")
	flag.Parse()

	db, err := sql.Open("mysql", fmt.Sprintf("root@/%s", *dbName))
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	stmtOut, err := db.Prepare("SELECT seq,  mid FROM member_information")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query()
	defer rows.Close()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var col1, col2 []byte
	for rows.Next() {
		err := rows.Scan(&col1, &col2)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		fmt.Println(string(col1), string(col2))
	}
}
