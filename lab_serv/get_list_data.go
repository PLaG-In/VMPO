package main

import (
	"encoding/json"
	"net/http"
)

func get_list_data(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("Secret")
	date := r.FormValue("Date")
	if check_session(key){
		answer := list_data(date)
		PrintToScreen(w, answer)
	}
	else {
		authAndRegFailed := AuthAndRegFAIL{403, "Неправильный ключ"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		PrintToScreen(w, js)
	}
}

//Для юнит-тестов
func list_data(login string, pass string) []byte {
	//Поиск в бд НЕ ГОТОВО
	rows, err := GetAnswer("")
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
