package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func Inc_Test(t int) int {
	return t + 1
}

func Inc_Failed_Test(f int) int {
	return f + 1
}

func testing(w http.ResponseWriter, r *http.Request) {
	failed_test_count := 0
	total_test := 0
	//1 выход из сети, по несуществующей сессии
	sign_out_data := Success{}
	answer := sign_out_test("asdas2312", "1")
	bytes := []byte(answer)
	json.Unmarshal(bytes, &sign_out_data)
	total_test = Inc_Test(total_test)
	if sign_out_data.Code == 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
	}
	//2 регистрация уже существующего пользователя
	reg_data := AuthAndRegOK{}
	answer = get_reg("12325sds", "345")
	bytes = []byte(answer)
	json.Unmarshal(bytes, &reg_data)
	total_test = Inc_Test(total_test)
	if reg_data.Code == 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
	}
	//3 авторизация не существующего пользователя
	auth_data := AuthAndRegOK{}
	answer = get_login("SergoGarage", "ssdd")
	bytes = []byte(answer)
	json.Unmarshal(bytes, &auth_data)
	total_test = Inc_Test(total_test)
	if auth_data.Code == 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
	}
	//5 получение описания по

	//Юнит-тесты 1 регистрируем нового пользователя
	reg_data = AuthAndRegOK{}
	answer = get_reg("unit", "test")
	bytes = []byte(answer)
	json.Unmarshal(bytes, &reg_data)
	total_test = Inc_Test(total_test)
	if reg_data.Code != 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
	}
	//7 получение данных по другому id
	list_data_struct := TaskList{}
	answer = list_data("04.02.2018", "1")
	bytes = []byte(answer)
	json.Unmarshal(bytes, &list_data_struct)
	total_test = Inc_Test(total_test)
	if list_data_struct.Code == 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
	}
	//2 выходим из сети
	sign_out_data = Success{}
	answer = sign_out_test(reg_data.SecretCode, strconv.Itoa(reg_data.User_id))
	bytes = []byte(answer)
	json.Unmarshal(bytes, &sign_out_data)
	total_test = Inc_Test(total_test)
	if sign_out_data.Code != 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
	}
	//3 авторизация
	auth_data = AuthAndRegOK{}
	answer = get_login("unit", "test")
	bytes = []byte(answer)
	json.Unmarshal(bytes, &auth_data)
	total_test = Inc_Test(total_test)
	if auth_data.Code != 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
	}
	//4 загружаем список с данными 4ого февраля
	list_data_struct = TaskList{}
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
	//6 добавление записи
	add_data := Success{}
	answer = append_data("5", "test", "test des", "05.02.2018", "22", "0")
	bytes = []byte(answer)
	json.Unmarshal(bytes, &add_data)
	total_test = Inc_Test(total_test)
	if add_data.Code != 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
	}
	//7 редактирование записи
	//8 удаление записи
	//9 стоп таймер
	//10 выходим из сети
	sign_out_data = Success{}
	answer = sign_out_test(auth_data.SecretCode, strconv.Itoa(auth_data.User_id))
	bytes = []byte(answer)
	json.Unmarshal(bytes, &sign_out_data)
	total_test = Inc_Test(total_test)
	if sign_out_data.Code != 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
	}
	w.Write([]byte("Total test - " + strconv.Itoa(total_test) + ". Failed test - " + strconv.Itoa(failed_test_count)))
}
