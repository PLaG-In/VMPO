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
