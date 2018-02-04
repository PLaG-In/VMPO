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
	http.HandleFunc("/search", search)
	http.HandleFunc("/sign_out", sign_out)
	/*http.HandleFunc("/timer_stop", timer_stop)*/
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

func Inc_Test(t int) int {
	return t + 1
}

func Inc_Failed_Test(f int) int {
	return f + 1
}

func testing(w http.ResponseWriter, r *http.Request) {
	failed_test_count := 0
	total_test := 0
	//Юнит-тесты 1 регистрируем нового пользователя
	reg_data := AuthAndRegOK{}
	answer := get_reg("unit", "test")
	bytes := []byte(answer)
	json.Unmarshal(bytes, &reg_data)
	total_test = Inc_Test(total_test)
	if reg_data.Code != 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
	}
	//2 выходим из сети
	sign_out_data := Success{}
	answer = sign_out_test(reg_data.SecretCode)
	bytes = []byte(answer)
	json.Unmarshal(bytes, &sign_out_data)
	total_test = Inc_Test(total_test)
	if sign_out_data.Code != 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
	}
	//3 авторизация
	auth_data := AuthAndRegOK{}
	answer = get_login("unit", "test")
	bytes = []byte(answer)
	json.Unmarshal(bytes, &auth_data)
	total_test = Inc_Test(total_test)
	if auth_data.Code != 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
	}
	//4 загружаем список с данными 4ого февраля
	list_data_struct := TaskList{}
	answer = list_data("04.02.2018", "5")
	bytes = []byte(answer)
	json.Unmarshal(bytes, &list_data_struct)
	total_test = Inc_Test(total_test)
	if list_data_struct.Code != 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
	}
	//5 просмотр описания по id задачи
	des_data_struct := FailAnswer{}
	answer = task_des("5", "5")
	bytes = []byte(answer)
	json.Unmarshal(bytes, &des_data_struct)
	total_test = Inc_Test(total_test)
	if des_data_struct.Code != 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
	}
	//6 выходим из сети
	sign_out_data = Success{}
	answer = sign_out_test(auth_data.SecretCode)
	bytes = []byte(answer)
	json.Unmarshal(bytes, &sign_out_data)
	total_test = Inc_Test(total_test)
	if sign_out_data.Code != 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
	}
	w.Write([]byte("Total test - " + strconv.Itoa(total_test) + ". Failed test - " + strconv.Itoa(failed_test_count)))
}
