// db.go
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "root"
	DB_NAME     = "postgres"
)

func SelectDB(input string) (*sql.Rows, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	fmt.Println("# Querying")
	rows, err := db.Query(input)
	checkErr(err)

	return rows, err
}

func UpdateDB(input string) error {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()
	fmt.Println("# Inserting values")

	var lastInsertId int
	_ = db.QueryRow(input).Scan(&lastInsertId)
	//checkErr(err)

	return err
}

func DeleteFromDB(input string) error {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	var lastInsertId int
	err = db.QueryRow(input).Scan(&lastInsertId)
	checkErr(err)

	fmt.Println("# Deleting")
	_, err = db.Prepare(input)
	checkErr(err)

	//	res, err := stmt.Exec(lastInsertId)
	//	checkErr(err)

	//	affect, err := res.RowsAffected()
	//	checkErr(err)

	return err
}

//func GetAnswer(input string) (*sql.Rows, error) {
//	db, err := sql.Open("mysql", "root:root@/mydb")
//	checkErr(err)
//	rows, err := db.Query(input)
//	checkErr(err)
//	return rows, err
//}

//func Update_DB(input string) error {
//	db, err := sql.Open("mysql", "root:root@/mydb")
//	checkErr(err)
//	stmt, err := db.Prepare(input)
//	checkErr(err)
//	_, err = stmt.Exec()
//	checkErr(err)
//	return err
//}
