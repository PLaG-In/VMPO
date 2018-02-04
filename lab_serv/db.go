// db.go
package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetAnswer(input string) (*sql.Rows, error) {
	db, err := sql.Open("mysql", "root:root@/mydb")
	checkErr(err)
	rows, err := db.Query(input)
	checkErr(err)
	return rows, err
}

func Update_DB(input string) error {
	db, err := sql.Open("mysql", "root:root@/mydb")
	checkErr(err)
	stmt, err := db.Prepare(input)
	checkErr(err)
	_, err = stmt.Exec()
	checkErr(err)
	return err
}
