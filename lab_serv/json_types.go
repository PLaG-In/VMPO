package main

type AuthAndRegOK struct {
	Code       int
	SecretCode string
	username   int
}

type AuthAndRegFAIL struct {
	Code        int
	Description string
}

type Task struct {
	Name        string
	Description string
}

type TaskList struct {
	task Task
}
