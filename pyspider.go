package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/darrenxyli/greedy/database/postgre"
	"github.com/darrenxyli/greedy/libs/result"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// Content is db content
type Content struct {
	URL      string
	Duration string
	Site     string
	Img      string
	Title    string
}

func main() {

	resultDB := postgre.NewResultDB(
		"amazon.cbtwp3cmfmsx.us-west-2.rds.amazonaws.com",
		5432,
		"ocean",
		"darrenxyli",
		"2jaqx97j",
		[]string{"porn"})

	db, err := sql.Open("mysql", "root:2jaqx97j@tcp(104.236.34.46:3306)/resultdb")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM pornhub")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Fetch rows
	for rows.Next() {
		var taskid string
		var url string
		var content string
		var updatetime string

		// get RawBytes from data
		err = rows.Scan(&taskid, &url, &content, &updatetime)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		var s Content
		json.Unmarshal([]byte(content), &s)
		url = s.URL
		dur, _ := strconv.Atoi(s.Duration)
		duration := uint(dur)
		site := s.Site
		img := s.Img
		title := s.Title
		fmt.Println(taskid, ": ", site, ": ", url, ": ", img, ": ", title, ": ", duration)
		resItem := result.NewResult(url, "porn", duration, site, img, title)
		resultDB.Insert(resItem)
		stmt, err := db.Prepare("DELETE FROM pornhub WHERE taskid=?")
		checkErr(err)
		stmt.Exec(taskid)
		checkErr(err)
		fmt.Println("-----------------------------------")
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
