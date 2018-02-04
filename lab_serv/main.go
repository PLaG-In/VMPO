package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	http.HandleFunc("/get_task_des", get_task_des)
	http.HandleFunc("/add_task", add_task)
	http.HandleFunc("/remove_task", remove_task)
	http.HandleFunc("/edit_task", edit_task)
	/*http.HandleFunc("/search", search)
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
		authAndRegFailed := FailAnswer{500, "Серверная ошибка"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		PrintToScreen(w, js)
	}
	PrintToScreen(w, js)
}

func testing(w http.ResponseWriter, r *http.Request) {
	failed_test_count := 0
	total_test := 0
	//Юнит-тесты 1 регистрируем нового пользователя
	reg_data := AuthAndRegOK{}
	answer := get_reg("unit", "test")
	bytes := []byte(answer)
	json.Unmarshal(bytes, &reg_data)
	if reg_data.Code == 200 {
		total_test = total_test + 1
	} else {
		total_test = total_test + 1
		failed_test_count = failed_test_count + 1
	}
	//2 выходим из сети
	//3 авторизация
	//4
	//5
	//6 выходим из сети
	w.Write([]byte("Total test - " + strconv.Itoa(total_test) + ". Failed test - " + strconv.Itoa(failed_test_count)))
}
