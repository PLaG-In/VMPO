// auth
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func auth(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	pass := r.FormValue("password")

	fmt.Println("login: " + login + " pass: " + pass)

	answer := get_login(login, pass)
	PrintToScreen(w, answer)
}

//Для юнит-тестов
func get_login(login string, pass string) []byte {
	//Поиск в бд
	rows, err := SelectDB("SELECT idusers FROM users WHERE login='" + login + "' AND password='" + pass + "'")
	if err != nil {
		authAndRegFailed := FailAnswer{500, "Серверная ошибка"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		return js
	}
	for rows.Next() {
		var username int
		err := rows.Scan(&username)
		if err != nil {
			authAndRegFailed := FailAnswer{500, "Серверная ошибка"}
			js, err := json.Marshal(authAndRegFailed)
			checkErr(err)
			return js
		}
		rows, err = SelectDB("SELECT idusers FROM users WHERE login = '" + login + "'")
		if err != nil {
			authAndRegFailed := FailAnswer{500, "Серверная ошибка"}
			js, err := json.Marshal(authAndRegFailed)
			checkErr(err)
			return js
		}
		uid := 0
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
	authAndRegFailed := FailAnswer{403, "Неправильный пароль"}
	js, err := json.Marshal(authAndRegFailed)
	checkErr(err)
	return js
}
