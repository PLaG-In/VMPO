// auth
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func remove_task(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("Secret")
	id_task := r.FormValue("id_task")
	id_user := r.FormValue("id_user")
	date := r.FormValue("date")
	fmt.Println(key)
	if check_session(key) {
		answer := task_delete(id_task, id_user, date)
		PrintToScreen(w, answer)
	} else {
		authAndRegFailed := FailAnswer{403, "Неправильный ключ"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		PrintToScreen(w, js)
	}
}

//Для юнит-тестов
func task_delete(id_task string, id_user string, date string) []byte {
	//Поиск в бд
	//Необходим будет фикс после реализации бд
	err := Update_DB("DELETE FROM mydb.task WHERE (task.id_user = " + id_user + ") AND (task.date = \"" + date + "\") AND (task.id_task = \"" + id_task + "\")")
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
