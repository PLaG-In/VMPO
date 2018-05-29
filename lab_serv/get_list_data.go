package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func get_list_data(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("secret")
	date := r.FormValue("date")
	id_user := r.FormValue("user_id")
	if check_session(key, id_user) {
		answer := list_data(date, id_user)
		PrintToScreen(w, answer)
	} else {
		authAndRegFailed := FailAnswer{403, "Неправильный ключ"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		PrintToScreen(w, js)
	}
}

//Для юнит-тестов
func list_data(date string, user string) []byte {
	//Поиск в бд НЕ ГОТОВО
	fmt.Println("SELECT * FROM task WHERE (task.iduser = " +
		user + ") AND (task.date = '" + date + "');")
	rows, err := SelectDB("SELECT * FROM task WHERE (task.iduser = " +
		user + ") AND (task.date = '" + date + "');")
	var i = 0
	for rows.Next() {
		i = i + 1
	}
	rows.Close()

	if i == 0 {
		authAndRegFailed := FailAnswer{201, "Нет данных по дате"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		return js
	}

	rows, err = SelectDB("SELECT idtask, name, time, des FROM task WHERE (task.iduser = " +
		user + ") AND (task.date = '" + date + "');")

	var tasks []Task = make([]Task, i)
	var counter = 0
	for rows.Next() {
		var uid int
		var name string
		var des string
		var time_task string
		err := rows.Scan(&uid, &name, &des, &time_task)

		checkErr(err)
		tasks[counter] = Task{uid, name, des, time_task}
		counter = counter + 1
	}
	fmt.Println(i)
	authAndRegOK := TaskList{200, i, tasks}
	js, err := json.Marshal(authAndRegOK)
	if err != nil {
		authAndRegFailed := FailAnswer{500, "Серверная ошибка"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		return js
	}
	return js
}
