package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func get_list_data(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("Secret")
	date := r.FormValue("Date")
	if check_session(key) {
		answer := list_data(date)
		PrintToScreen(w, answer)
	} else {
		authAndRegFailed := FailAnswer{403, "Неправильный ключ"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		PrintToScreen(w, js)
	}
}

//Для юнит-тестов
func list_data(date string) []byte {
	//Поиск в бд НЕ ГОТОВО
	rows, err := GetAnswer("")
	i := checkCount(rows)

	var tasks []Task = make([]Task, i)
	var counter = 0
	for rows.Next() {
		var uid int
		var name string
		var des string
		var date string
		var time_task string
		err := rows.Scan(&uid, &name, &des, &date, &time_task)

		if err != nil {
			authAndRegFailed := FailAnswer{500, "Серверная ошибка"}
			js, err := json.Marshal(authAndRegFailed)
			checkErr(err)
			return js
		}

		tasks[counter] = Task{uid, name, des, date, time_task}
		counter = counter + 1
	}
	task_list := TaskList{200, i, tasks}
	js, err := json.Marshal(task_list)
	if err != nil {
		authAndRegFailed := FailAnswer{500, "Серверная ошибка"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		return js
	}
	return js
}

func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		checkErr(err)
	}
	return count
}
