// auth
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func get_task_des(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("secret")
	task_id := r.FormValue("task_id")
	user_id := r.FormValue("user_id")
	if check_session(key, user_id) {
		answer := task_des(task_id, user_id)
		PrintToScreen(w, answer)
	} else {
		authAndRegFailed := FailAnswer{403, "Неправильный ключ"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		PrintToScreen(w, js)
	}
}

//Для юнит-тестов
func task_des(id string, user string) []byte {
	//Поиск в бд
	fmt.Println("SELECT des FROM task WHERE (task.idtask = " + id + ")  AND (task.iduser = " + user + ");")
	rows, err := SelectDB("SELECT des FROM task WHERE (task.idtask = " + id + ")  AND (task.iduser = " + user + ");")
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
