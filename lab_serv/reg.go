// reg.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func reg(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	pass := r.FormValue("password")

	answer := get_reg(login, pass)
	PrintToScreen(w, answer)
}

func get_reg(login string, pass string) []byte {
	rows, err := SelectDB("SELECT idusers FROM users WHERE login = '" + login + "'")
	if err != nil {
		authAndRegFailed := FailAnswer{500, "Серверная ошибка"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		return js
	}
	var uid int
	for rows.Next() {
		//Error логин существует
		authAndRegFailed := FailAnswer{403, "Данный логин уже существует"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		return js
	}
	fmt.Println("INSERT INTO users (login, password) VALUES ('" + login + "', '" + pass + "');")
	//INSERT INTO users (login, password) VALUES ('" + login + "', '" + password + "');
	err = UpdateDB("INSERT INTO users (login, password) VALUES ('" + login + "', '" + pass + "');")

	//	if err != nil {
	//		authAndRegFailed := FailAnswer{500, "Серверная ошибка"}
	//		js, err := json.Marshal(authAndRegFailed)
	//		checkErr(err)
	//		return js
	//	}
	//JSON ответ успешной регистрации
	fmt.Println("SELECT idusers FROM users WHERE login = '" + login + "'")
	rows, err = SelectDB("SELECT idusers FROM users WHERE login = '" + login + "'")
	if err != nil {
		authAndRegFailed := FailAnswer{500, "Серверная ошибка"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		return js
	}
	for rows.Next() {
		var user_id int
		err := rows.Scan(&user_id)
		if err != nil {
			authAndRegFailed := FailAnswer{500, "Серверная ошибка"}
			js, err := json.Marshal(authAndRegFailed)
			checkErr(err)
			return js
		}
		uid = user_id
	}

	authAndRegOK := AuthAndRegOK{200, start_session(strconv.Itoa(uid)), uid}
	js, err := json.Marshal(authAndRegOK)
	if err != nil {
		authAndRegFailed := FailAnswer{500, "Серверная ошибка"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		return js
	}

	return js
}
