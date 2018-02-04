package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func sign_out(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("Secret")
	if check_session(key) {
		answer := sign_out_test(key)
		PrintToScreen(w, answer)
	} else {
		authAndRegFailed := FailAnswer{403, "Неправильный ключ"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		PrintToScreen(w, js)
	}
}

//Для юнит-тестов
func sign_out_test(key string) []byte {
	closeSession(key)
	result := Success{200}
	fmt.Println("LOG OUT")
	js, err := json.Marshal(result)
	checkErr(err)
	return js
}
