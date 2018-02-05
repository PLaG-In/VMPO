// auth
package main

import (
	"encoding/json"
	"net/http"
)

func add_task(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("Secret")
	id := r.FormValue("user_id")
	name := r.FormValue("name")
	data := r.FormValue("date")
	time := r.FormValue("time")
	priority := r.FormValue("priority")
	description := r.FormValue("description")
	if check_session(key, id) {
		answer := append_data(id, name, description, data, time, priority)
		PrintToScreen(w, answer)
	} else {
		authAndRegFailed := FailAnswer{403, "Неправильный ключ"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		PrintToScreen(w, js)
	}
}

//Для юнит-тестов
func append_data(user string, name string, des string, date string, time string, priority string) []byte {
	//Поиск в бд
	//fmt.Println("INSERT INTO mydb.task (id_user, name, des, date, time, priority) VALUES (" + user + ", \"" + name + "\", \"" + des + "\", \"" + date + "\", \"" + time + "\", \"" + priority + "\");")
	err := Update_DB("INSERT INTO mydb.task (id_user, name, des, date, time, priority) VALUES (\"" + user + "\", \"" + name + "\", \"" + des + "\", \"" + date + "\", \"" + time + "\", \"" + priority + "\");")
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
