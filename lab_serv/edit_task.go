// auth
package main

import (
	"encoding/json"
	"net/http"
)

func edit_task(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("Secret")
	id_task := r.FormValue("id_task")
	name := r.FormValue("name")
	data := r.FormValue("data")
	priority := r.FormValue("priority")
	description := r.FormValue("description")
	if check_session(key) {
		answer := task_update(id_task, name, data, priority, description)
		PrintToScreen(w, answer)
	} else {
		authAndRegFailed := FailAnswer{403, "Неправильный ключ"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		PrintToScreen(w, js)
	}
}

//Для юнит-тестов
func task_update(id_task string, name string, data string, priority string, description string) []byte {
	//Поиск в бд
	err := Update_DB("")
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
