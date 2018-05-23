// auth
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func timer_stop(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("secret")
	id_task := r.FormValue("task_id")
	id_user := r.FormValue("user_id")
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
	fmt.Println("UPDATE task SET time=" + time + " where (idtask = " + id_task + " AND iduser = " + id_user + ");")
	err := UpdateDB("UPDATE task SET time=" + time + " where (idtask = " + id_task + " AND iduser = " + id_user + ");")
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
