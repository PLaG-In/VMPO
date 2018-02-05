// auth
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func timer_stop(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("Secret")
	id_task := r.FormValue("id_task")
	id_user := r.FormValue("id_user")
	name := r.FormValue("name")
	time := r.FormValue("time")
	fmt.Println(key)
	if check_session(key, id_user) {
		answer := timer_unit(id_task, id_user, name, time)
		PrintToScreen(w, answer)
	} else {
		authAndRegFailed := FailAnswer{403, "Неправильный ключ"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		PrintToScreen(w, js)
	}
}

//Для юнит-тестов
func timer_unit(id_task string, id_user string, name string, time string) []byte {
	//Поиск в бд
	//Необходим будет фикс после реализации бд
	err := Update_DB("UPDATE task SET task.time=\"" + time + "\" where (task.idtask = " + id_task + " AND task.id_user = " + id_user + ")")
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
