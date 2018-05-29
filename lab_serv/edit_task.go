// auth
package main

import (
	"encoding/json"
	"net/http"
)

func edit_task(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("secret")
	id := r.FormValue("user_id")
	name := r.FormValue("name")
	//data := r.FormValue("date")
	time := r.FormValue("time")
	//priority := r.FormValue("priority")
	description := r.FormValue("description")
	//old_date := r.FormValue("old_date")
	task_id := r.FormValue("task_id")
	if check_session(key, id) {
		answer := task_update(id, name, description, task_id, time)
		PrintToScreen(w, answer)
	} else {
		authAndRegFailed := FailAnswer{403, "Неправильный ключ"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		PrintToScreen(w, js)
	}
}

//Для юнит-тестов
func task_update(user string, name string, des string, task_id string, time string) []byte {
	//Поиск в бд
	//fmt.Println("UPDATE task SET (name, des, time) = ('" + name + "', '" +
	//	des + "', '" + time + "') WHERE (iduser = " + user + ") AND (idtask = " + task_id + ");")
	err := UpdateDB("UPDATE task SET (name, des, time) = ('" + name + "', '" +
		des + "', '" + time + "') WHERE (iduser = " + user + ") AND (idtask = " + task_id + ");")

	//	if err != nil {
	//		authAndRegFailed := FailAnswer{500, "Серверная ошибка"}
	//		js, err := json.Marshal(authAndRegFailed)
	//		checkErr(err)
	//		return js
	//	}
	result := Success{200}
	js, err := json.Marshal(result)
	checkErr(err)
	return js
}
