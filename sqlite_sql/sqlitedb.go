package main

import (
	"database/sql" //golang 定義的SQL標準介面, 讓其它DATABASE Driver能用這個標準,進而使用致的方法
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3" //已有實做database/sql, 但需將CGO_ENABLED=1打開, 並有gcc compiler
)

func checkErr(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {
	const create string = `
	CREATE TABLE IF NOT EXISTS activities (
    id INTEGER NOT NULL PRIMARY KEY,
	time DATETIME NOT NULL,
	description TEXT
	);`

	db, err := sql.Open("sqlite3", "./sqlite.db")
	checkErr(err)
	defer db.Close()

	_, err = db.Exec(create)
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO activities VALUES(?,?,?);")
	checkErr(err)
	defer stmt.Close()
	res, err := stmt.Exec(1, "2023-03-27T16:56:12", "my sql test.")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println("Last Insert Id: ", id)

	rows, err := db.Query("select * from activities;")
	checkErr(err)

	for rows.Next() {
		var id int
		var time time.Time
		var description string
		err := rows.Scan(&id, &time, &description)
		checkErr(err)
		fmt.Println("Id: ", id)
		fmt.Println("time: ", time)
		fmt.Println("description: ", description)
	}
	defer rows.Close()

}
