// auth
package main

import (
	"encoding/json"
	"net/http"
)

func remove_task_all(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("Secret")
	id_user := r.FormValue("id_user")
	date := r.FormValue("date")
	if check_session(key) {
		answer := task_delete_all(id_user, date)
		PrintToScreen(w, answer)
	} else {
		authAndRegFailed := FailAnswer{403, "Неправильный ключ"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		PrintToScreen(w, js)
	}
}

//Для юнит-тестов
func task_delete_all(id_task string, date string) []byte {
	//Поиск в бд
	//Необходим будет фикс после реализации бд
	err := Update("")
	if err != nil {
		authAndRegFailed := FailAnswer{500, "Серверная ошибка"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		return js
	}
	result := Success{200}
	js, err := json.Marshal(result)
	checkErr(err)
	return js
}
