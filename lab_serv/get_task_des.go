// auth
package main

import (
	"encoding/json"
	"net/http"
)

func get_task_des(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("Secret")
	task_id := r.FormValue("id_task")
	if check_session(key) {
		answer := task_des(task_id)
		PrintToScreen(w, answer)
	} else {
		authAndRegFailed := FailAnswer{403, "Неправильный ключ"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		PrintToScreen(w, js)
	}
}

//Для юнит-тестов
func task_des(id string) []byte {
	//Поиск в бд
	rows, err := GetAnswer("")
	if err != nil {
		authAndRegFailed := FailAnswer{500, "Серверная ошибка"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		return js
	}
	for rows.Next() {
		var des string
		err := rows.Scan(&des)
		if err != nil {
			authAndRegFailed := FailAnswer{500, "Серверная ошибка"}
			js, err := json.Marshal(authAndRegFailed)
			checkErr(err)
			return js
		}
		authAndRegOK := FailAnswer{200, des}
		js, err := json.Marshal(authAndRegOK)
		if err != nil {
			authAndRegFailed := FailAnswer{500, "Серверная ошибка"}
			js, err := json.Marshal(authAndRegFailed)
			checkErr(err)
			return js
		}
		return js
	}
	authAndRegFailed := FailAnswer{403, "Задачи нет в базе"}
	js, err := json.Marshal(authAndRegFailed)
	checkErr(err)
	return js
}
