package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func PrintToScreen(w http.ResponseWriter, js []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	http.HandleFunc("/", checkState)
	http.HandleFunc("/auth", auth)
	http.HandleFunc("/reg", reg)
	http.HandleFunc("/get_list_data", get_list_data)
	/*http.HandleFunc("/get_task_des", get_task_des)
	http.HandleFunc("/add_task", add_task)
	http.HandleFunc("/remove_task", remove_task)
	http.HandleFunc("/remove_task_all", remove_task_all)
	http.HandleFunc("/edit_task", edit_task)
	http.HandleFunc("/search", search)
	http.HandleFunc("/sign_out", sign_out)
	http.HandleFunc("/timer_stop", timer_stop)*/
	http.HandleFunc("/testing", testing)
	http.ListenAndServe(":8080", nil)
}

//Ошибки
func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		//log.Fatal(err.Error())
	}
}

func checkState(w http.ResponseWriter, r *http.Request) {
	//Проверка состояния сервера
	authAndRegOK := AuthAndRegOK{200, "", 0}
	js, err := json.Marshal(authAndRegOK)
	if err != nil {
		authAndRegFailed := AuthAndRegFAIL{500, "Серверная ошибка"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		PrintToScreen(w, js)
	}
	PrintToScreen(w, js)
}

func testing(w http.ResponseWriter, r *http.Request) {
	//Юнит-тесты
}
