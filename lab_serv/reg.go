// reg.go
package main

import (
	"encoding/json"
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
	rows, err := GetAnswer("SELECT idusers FROM users WHERE login = \"" + login + "\"")
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

	err = Update_DB("INSERT users SET login=\"" + login +
		"\",password=\"" + pass + "\"")
	if err != nil {
		authAndRegFailed := FailAnswer{500, "Серверная ошибка"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		return js
	}
	//JSON ответ успешной регистрации
	rows, err = GetAnswer("SELECT idusers FROM users WHERE login = \"" + login + "\"")
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
