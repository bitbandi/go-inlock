package main

import (
	"fmt"

	"github.com/bitbandi/go-inlock"
)

const (
	ACCESS_TOKEN  = ""
	REFRESH_TOKEN = ""
)

func main() {
	// inlock client
	inlock := inlock.New(ACCESS_TOKEN, REFRESH_TOKEN)

	// Login
	tokens, _ := inlock.Login("username", "password")
	fmt.Println(tokens)
}
