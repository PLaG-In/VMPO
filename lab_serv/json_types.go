package main

type AuthAndRegOK struct {
	Code       int
	SecretCode string
	User_id    int
}

type FailAnswer struct {
	Code        int
	Description string
}

type Task struct {
	Id_task int
	Name    string
	Time    string
}

type TaskList struct {
	Code   int
	Length int
	Task   []Task
}

type Success struct {
	Code int
}

type AppendData struct {
	Code    int
	Id_task string
}
