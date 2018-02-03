package main

type AuthAndRegOK struct {
	Code       int
	SecretCode string
	username   int
}

type FailAnswer struct {
	Code        int
	Description string
}

type Task struct {
	Id_task     int
	Name        string
	Description string
	Date        string //????
	Time        string
}

type TaskList struct {
	Code   int
	Length int
	task   []Task
}
