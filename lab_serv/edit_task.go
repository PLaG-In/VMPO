// auth
package main

import (
	"encoding/json"
	"net/http"
)

func edit_task(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("Secret")
	id := r.FormValue("user_id")
	name := r.FormValue("name")
	data := r.FormValue("date")
	time := r.FormValue("time")
	priority := r.FormValue("priority")
	description := r.FormValue("description")
	old_date := r.FormValue("old_date")
	task_id := r.FormValue("task_id")
	if check_session(key) {
		answer := task_update(id, name, description, data, time, priority, old_date, task_id)
		PrintToScreen(w, answer)
	} else {
		authAndRegFailed := FailAnswer{403, "Неправильный ключ"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		PrintToScreen(w, js)
	}
}

//Для юнит-тестов
func task_update(user string, name string, des string, date string, time string, priority string, old_date string, task_id string) []byte {
	//Поиск в бд
	err := Update_DB("UPDATE mydb.task SET task.name = \"" + name + "\", task.des = " + des + ", task.time = \"" + time + "\", task.date = \"" + date + "\", task.priority = \"" + priority + "\" WHERE (task.id_user = " + user + ") AND (task.date = \"" + old_date + "\") AND (task.id_task = \"" + task_id + "\")")
	task_delete(task_id, user, old_date)

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
