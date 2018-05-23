// auth
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func add_task(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("secret")
	id := r.FormValue("user_id")
	name := r.FormValue("name")
	data := r.FormValue("date")
	time := r.FormValue("time")
	//priority := r.FormValue("priority")
	description := r.FormValue("description")
	if check_session(key, id) {
		answer := append_data(id, name, description, data, time)
		PrintToScreen(w, answer)
	} else {
		authAndRegFailed := FailAnswer{403, "Неправильный ключ"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		PrintToScreen(w, js)
	}
}

//Для юнит-тестов
func append_data(user string, name string, des string, date string, time string) []byte {
	//Поиск в бд
	fmt.Println("INSERT INTO task (iduser, name, des, date, time) VALUES (" + user + ", '" + name + "', '" +
		des + "', '" + date + "', " + time + ");")
	err := UpdateDB("INSERT INTO task (iduser, name, des, date, time) VALUES (" + user + ", '" + name + "', '" +
		des + "', '" + date + "', " + time + ");")
	/*if err != nil {
		authAndRegFailed := FailAnswer{500, "Серверная ошибка"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		return js
	}*/
	fmt.Println("SELECT task.idtask from task WHERE (task.name='" + name + "' and task.des='" + des + "' and task.date='" + date + "');")
	answer, err := SelectDB("SELECT task.idtask from task WHERE (task.name='" + name + "' and task.des='" + des + "' and task.date='" + date + "');")
	var uid string
	for answer.Next() {
		//Error логин существует
		err = answer.Scan(&uid)
		checkErr(err)
	}
	result := AppendData{200, uid}
	js, err := json.Marshal(result)
	checkErr(err)
	return js
}
