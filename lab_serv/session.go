package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

var secret_range = 25
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func secret() string {
	b := make([]rune, secret_range)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func exists(path string, id string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	} else {
		var file, err = os.Open(path)
		checkErr(err)
		b, err := ioutil.ReadFile(path) // just pass the file name
		if err != nil {
			fmt.Print(err)
		}
		if string(b) == id {
			defer file.Close()
			return true, nil
		}
		defer file.Close()
		return false, nil
	}

	return true, err
}

func createFile(path string, id string) {
	// detect if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		checkErr(err)
		fmt.Fprintf(file, id)
		defer file.Close()
	}

	fmt.Println("==> done creating file", path)
}

func closeSession(path string) {
	// detect if file exists
	path = "./Session/" + path + ".txt"
	var _, err = os.Stat(path)

	checkErr(err)

	// create file if not exists
	err = os.Remove(path)
	checkErr(err)

	fmt.Println("==> delete file", path)
}

func check_session(key string, id string) bool {
	answer, err := exists("./Session/"+key+".txt", id)
	checkErr(err)
	return answer
}

func start_session(id string) string {
	rand.Seed(time.Now().UnixNano())
	random := secret()
	key, err := exists("./Session/"+random+".txt", id)
	checkErr(err)
	for key {
		key, err = exists("./Session/"+random+".txt", id)
		checkErr(err)
		fmt.Println("123")
	}
	createFile("./Session/"+random+".txt", id)
	return random
}
