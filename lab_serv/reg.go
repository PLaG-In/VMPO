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
	rows, err := GetAnswer("SELECT id FROM users WHERE login = \"" + login + "\"")
	if err != nil {
		authAndRegFailed := FailAnswer{500, "Серверная ошибка"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		return js
	}
	var uid int
	for rows.Next() {
		//Error логин существует
		authAndRegFailed := FailAnswer{500, "Серверная ошибка"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		return js
	}

	err = Update("INSERT users SET idusers=\"" + strconv.Itoa(uid+1) + "\", login=\"" + login +
		"\",pass=\"" + pass + "\"")
	if err != nil {
		authAndRegFailed := FailAnswer{500, "Серверная ошибка"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		return js
	}
	//JSON ответ успешной регистрации
	authAndRegOK := AuthAndRegOK{200, start_session(), uid + 1}
	js, err := json.Marshal(authAndRegOK)
	if err != nil {
		authAndRegFailed := FailAnswer{500, "Серверная ошибка"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		return js
	}

	return js
}
