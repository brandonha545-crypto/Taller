package main

import (
	"fmt"
	"strings"
)

func main() {

	username := "Juan342"
	email := "mail@gmail.com"
	password := "12345678"
	age := 18

	var errors []string

	if len(username) >= 4 && len(username) <= 20 {
		fmt.Println("Username correcto")
	} else {
		errors = append(errors, "Username incorrecto")
	}

	if strings.Contains(email, "@") && strings.Contains(email, ".") {
		fmt.Println("Email correcto")
	} else {
		errors = append(errors, "Email incorrecto")
	}

	if len(password) >= 8 {
		fmt.Println("Password correcto")
	} else {
		errors = append(errors, "Password incorrecto")
	}

	if age >= 18 {
		fmt.Println("Age correcto")
	} else {
		errors = append(errors, "Age incorrecto")
	}

	if len(errors) == 0 {
		fmt.Println("Registro exitoso")
	} else {
		fmt.Println("Error:")
		fmt.Println(errors)
	}
}
