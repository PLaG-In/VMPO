package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func search(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("Secret")
	search := r.FormValue("search_string")
	id_user := r.FormValue("id_user")
	if check_session(key) {
		answer := get_search(search, id_user)
		PrintToScreen(w, answer)
	} else {
		authAndRegFailed := FailAnswer{403, "Неправильный ключ"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		PrintToScreen(w, js)
	}
}

//Для юнит-тестов
func get_search(search string, user string) []byte {
	//Поиск в бд НЕ ГОТОВО
	rows, err := GetAnswer("SELECT * FROM mydb.users, mydb.task WHERE (task.id_user= " + user + ") and (task.des RLIKE \"" + search + "\" ORDER BY LENGTH(\"" + search + "\"))")
	var i = 0
	for rows.Next() {
		i = i + 1
	}
	rows.Close()

	if i == 0 {
		authAndRegFailed := FailAnswer{402, "Нет данных по дате"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		return js
	}

	rows, err = GetAnswer("SELECT idtask, name, time, priority FROM mydb.users, mydb.task WHERE (task.id_user= " + user + ") and (task.des LIKE \"" + search + "\" + '%' ORDER BY LENGTH(\"" + search + "\") DESC LIMIT 1)")

	var tasks []Task = make([]Task, i)
	var counter = 0
	for rows.Next() {
		var uid int
		var name string
		var time_task string
		var priority string
		err := rows.Scan(&uid, &name, &time_task, &priority)

		checkErr(err)
		tasks[counter] = Task{uid, name, time_task, priority}
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

func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		checkErr(err)
	}
	return count
}
