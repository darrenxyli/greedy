package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:2jaqx97j@tcp(104.236.34.46:3306)/resultdb")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM pornhub LIMIT 5")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for _, v := range columns {
		fmt.Println(v)
	}

	// Fetch rows
	for rows.Next() {
		var taskid string
		var url string
		var result string
		var updatetime string

		// get RawBytes from data
		err = rows.Scan(&taskid, &url, &result, &updatetime)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		fmt.Println(taskid, ": ", url, ": ", result, ": ", updatetime)
		fmt.Println("-----------------------------------")
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
