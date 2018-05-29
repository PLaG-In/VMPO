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
		w.Write([]byte("1 failed"))
	}
	//2 регистрация уже существующего пользователя
	reg_data := AuthAndRegOK{}
	answer = get_reg("12325sds", "345")
	bytes = []byte(answer)
	json.Unmarshal(bytes, &reg_data)
	total_test = Inc_Test(total_test)
	if reg_data.Code == 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("2 failed"))
	}
	//	//3 авторизация не существующего пользователя
	auth_data := AuthAndRegOK{}
	answer = get_login("SergoGarage", "ssdd")
	bytes = []byte(answer)
	json.Unmarshal(bytes, &auth_data)
	total_test = Inc_Test(total_test)
	if auth_data.Code == 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("3 failed"))
	}
	//4.1 регистрируем нового пользователя
	//	reg_data = AuthAndRegOK{}
	//	answer = get_reg("unit", "test")
	//	bytes = []byte(answer)
	//	json.Unmarshal(bytes, &reg_data)
	//	total_test = Inc_Test(total_test)
	//	if reg_data.Code != 200 {
	//		failed_test_count = Inc_Failed_Test(failed_test_count)
	//	}
	//4.2 авторизация пользователя
	reg_data = AuthAndRegOK{}
	answer = get_login("unit", "test")
	bytes = []byte(answer)
	json.Unmarshal(bytes, &reg_data)
	total_test = Inc_Test(total_test)
	if reg_data.Code != 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("4.2 failed"))
	}
	//5 выход
	sign_out_data = Success{}
	answer = sign_out_test(reg_data.SecretCode, strconv.Itoa(reg_data.User_id))
	bytes = []byte(answer)
	json.Unmarshal(bytes, &sign_out_data)
	total_test = Inc_Test(total_test)
	if sign_out_data.Code != 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("5 failed"))
	}
	//5.1 авторизация пользователя
	reg_data = AuthAndRegOK{}
	answer = get_login("unit", "test")
	bytes = []byte(answer)
	json.Unmarshal(bytes, &reg_data)
	//6  получение описания по задаче с другим id
	des_data_struct := FailAnswer{}
	answer = task_des("123", "1")
	bytes = []byte(answer)
	json.Unmarshal(bytes, &des_data_struct)
	total_test = Inc_Test(total_test)
	if des_data_struct.Code == 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("6 failed"))
	}
	//7 получение данных описания с невалидными полями
	des_data_struct = FailAnswer{}
	answer = task_des("фывф", "фыфаы")
	bytes = []byte(answer)
	json.Unmarshal(bytes, &des_data_struct)
	total_test = Inc_Test(total_test)
	if des_data_struct.Code == 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("7 failed"))
	}
	//8 получение данных по другому id
	list_data_struct := TaskList{}
	answer = list_data("2018-05-23", "12")
	bytes = []byte(answer)
	json.Unmarshal(bytes, &list_data_struct)
	total_test = Inc_Test(total_test)
	if list_data_struct.Code == 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("8 failed"))
	}
	//9 загружаем список с данными 4ого февраля (пустой список)
	list_data_struct = TaskList{}
	answer = list_data("02.04.2018", strconv.Itoa(reg_data.User_id))
	bytes = []byte(answer)
	json.Unmarshal(bytes, &list_data_struct)
	total_test = Inc_Test(total_test)
	if list_data_struct.Code != 201 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("9 failed"))
	}
	//10 добавить запись на 5 февраля
	add_data := AppendData{}
	answer0 := append_data(strconv.Itoa(reg_data.User_id), "test", "test des", "02.05.2018", "00:00:00")
	bytes = []byte(answer0)
	json.Unmarshal(bytes, &add_data)
	total_test = Inc_Test(total_test)
	if add_data.Code != 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("10 failed"))
	}
	//11, 12, 13 загружаем список с данными 5ого февраля (код, элемент = 1, имя элемента = test)
	list_data_struct = TaskList{}
	answer = list_data("02.05.2018", strconv.Itoa(reg_data.User_id))
	bytes = []byte(answer)
	json.Unmarshal(bytes, &list_data_struct)
	total_test = Inc_Test(total_test)
	if list_data_struct.Code != 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("11 failed"))
	}
	total_test = Inc_Test(total_test)
	if list_data_struct.Length != 1 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("12 failed"))
	}
	total_test = Inc_Test(total_test)
	taskData := Task{}
	taskData = list_data_struct.Task[0]
	if taskData.Name != "test" {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("13 failed"))
	}
	//еще 5
	add_data1 := AppendData{}
	add_data2 := AppendData{}
	add_data3 := AppendData{}
	add_data4 := AppendData{}
	add_data5 := AppendData{}
	answer1 := append_data(strconv.Itoa(reg_data.User_id), "test1", "test des", "02.05.2018", "00:00:00")
	bytes = []byte(answer1)
	json.Unmarshal(bytes, &add_data1)
	answer2 := append_data(strconv.Itoa(reg_data.User_id), "test2", "test des", "02.05.2018", "00:00:00")
	bytes = []byte(answer2)
	json.Unmarshal(bytes, &add_data2)
	answer3 := append_data(strconv.Itoa(reg_data.User_id), "test3", "test des", "02.05.2018", "00:00:00")
	bytes = []byte(answer3)
	json.Unmarshal(bytes, &add_data3)
	answer4 := append_data(strconv.Itoa(reg_data.User_id), "test4", "test des", "02.05.2018", "00:00:00")
	bytes = []byte(answer4)
	json.Unmarshal(bytes, &add_data4)
	answer5 := append_data(strconv.Itoa(reg_data.User_id), "test5", "test des", "02.05.2018", "00:00:00")
	bytes = []byte(answer5)
	json.Unmarshal(bytes, &add_data5)
	//14 количество записей - 6
	answer = list_data("02.05.2018", strconv.Itoa(reg_data.User_id))
	bytes = []byte(answer)
	json.Unmarshal(bytes, &list_data_struct)
	total_test = Inc_Test(total_test)
	if list_data_struct.Length != 6 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("14 failed"))
	}
	//15 удалить запись test5
	tastDelete := Success{}
	answer = task_delete(add_data5.Id_task, strconv.Itoa(reg_data.User_id))
	bytes = []byte(answer)
	json.Unmarshal(bytes, &tastDelete)
	total_test = Inc_Test(total_test)
	if tastDelete.Code != 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("15 failed"))
	}
	//16 количество записей - 5
	answer = list_data("02.05.2018", strconv.Itoa(reg_data.User_id))
	bytes = []byte(answer)
	json.Unmarshal(bytes, &list_data_struct)
	total_test = Inc_Test(total_test)
	if list_data_struct.Length != 5 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("16 failed"))
	}
	//17 редактирование записи test4
	taskEdit := Success{}
	answer = task_update(strconv.Itoa(reg_data.User_id), "TASK5", "none", add_data4.Id_task, "00:00:00") //(add_data5.Id_task, strconv.Itoa(reg_data.User_id))
	bytes = []byte(answer)
	json.Unmarshal(bytes, &taskEdit)
	total_test = Inc_Test(total_test)
	if taskEdit.Code != 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("17 failed"))
	}
	//18 количество записей - 5
	answer = list_data("02.05.2018", strconv.Itoa(reg_data.User_id))
	bytes = []byte(answer)
	json.Unmarshal(bytes, &list_data_struct)
	total_test = Inc_Test(total_test)
	if list_data_struct.Length != 5 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("18 failed"))
	}
	//19 проверить имя
	taskData = Task{}
	taskData = list_data_struct.Task[4]
	total_test = Inc_Test(total_test)
	if taskData.Name != "TASK5" {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("19 failed"))
	}
	//20 проверить описание
	total_test = Inc_Test(total_test)
	if taskData.Time != "none" {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("20 failed " + taskData.Description))
	}
	//21 повторно удалить task5
	tastDelete = Success{}
	answer = task_delete(add_data5.Id_task, strconv.Itoa(reg_data.User_id))
	bytes = []byte(answer)
	total_test = Inc_Test(total_test)
	json.Unmarshal(bytes, &tastDelete)
	if tastDelete.Code != 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("21 failed"))
	}
	//удаляем все, кроме одного
	answer = task_delete(add_data4.Id_task, strconv.Itoa(reg_data.User_id))
	bytes = []byte(answer)
	answer = task_delete(add_data3.Id_task, strconv.Itoa(reg_data.User_id))
	bytes = []byte(answer)
	answer = task_delete(add_data2.Id_task, strconv.Itoa(reg_data.User_id))
	bytes = []byte(answer)
	answer = task_delete(add_data1.Id_task, strconv.Itoa(reg_data.User_id))
	bytes = []byte(answer)
	//22 количество записей - 1
	answer = list_data("02.05.2018", strconv.Itoa(reg_data.User_id))
	bytes = []byte(answer)
	json.Unmarshal(bytes, &list_data_struct)
	total_test = Inc_Test(total_test)
	if list_data_struct.Length != 1 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("22 failed"))
	}
	//23 отредиктировать запись task (другая дата)
	taskEdit = Success{}
	answer = task_update(strconv.Itoa(reg_data.User_id), "task", "none", add_data.Id_task, "00:00:00") //(add_data5.Id_task, strconv.Itoa(reg_data.User_id))
	bytes = []byte(answer)
	json.Unmarshal(bytes, &tastDelete)
	total_test = Inc_Test(total_test)
	if tastDelete.Code != 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("23 failed"))
	}
	//удалить запись
	answer = task_delete(add_data.Id_task, strconv.Itoa(reg_data.User_id))
	bytes = []byte(answer)
	//выйти
	answer = sign_out_test(reg_data.SecretCode, strconv.Itoa(reg_data.User_id))
	bytes = []byte(answer)
	w.Write([]byte("Total test - " + strconv.Itoa(total_test) + ". Failed test - " + strconv.Itoa(failed_test_count)))
}
