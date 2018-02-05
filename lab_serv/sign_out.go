package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func sign_out(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("Secret")
	id := r.FormValue("user_id")
	answer := sign_out_test(key, id)
	PrintToScreen(w, answer)
}

//Для юнит-тестов
func sign_out_test(key string, id string) []byte {
	if check_session(key, id) {
		closeSession(key)
		result := Success{200}
		fmt.Println("LOG OUT")
		js, err := json.Marshal(result)
		checkErr(err)
		return js
	} else {
		authAndRegFailed := FailAnswer{403, "Неправильный ключ"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		return js
	}
}
