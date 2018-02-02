// auth
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	rows, err := GetAnswer("SELECT idusers FROM users WHERE login=\"" + login + "\" AND pass=\"" + pass + "\"")
	if err != nil {
		authAndRegFailed := AuthAndRegFAIL{500, "Серверная ошибка"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		return js
	}
	for rows.Next() {
		var username int
		err := rows.Scan(&username)
		if err != nil {
			authAndRegFailed := AuthAndRegFAIL{500, "Серверная ошибка"}
			js, err := json.Marshal(authAndRegFailed)
			checkErr(err)
			return js
		}
		authAndRegOK := AuthAndRegOK{200, start_session(), username}
		js, err := json.Marshal(authAndRegOK)
		if err != nil {
			authAndRegFailed := AuthAndRegFAIL{500, "Серверная ошибка"}
			js, err := json.Marshal(authAndRegFailed)
			checkErr(err)
			return js
		}
		return js
	}
	authAndRegFailed := AuthAndRegFAIL{403, "Неправильный пароль"}
	js, err := json.Marshal(authAndRegFailed)
	checkErr(err)
	return js
}
