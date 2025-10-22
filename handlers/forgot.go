package handlers

import (
	"Golang4/models"
	"bufio"
	"fmt"
)

func ForgotPassword(service models.AuthService, scanner *bufio.Scanner) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("\nError: %v\n", r)
			waitEnter(scanner)
		}
	}()

	defer waitEnter(scanner)

	fmt.Println("\n---- Forgot Password ----")
	fmt.Println()
	email := getInput("Enter your email: ", scanner)
	password := getInput("Enter a strong password: ", scanner)
	confirmPassword := getInput("Confirm your password: ", scanner)

	if password != confirmPassword {
		panic("Password doesn't match!")
	}

	err := service.ResetPassword(email, password)
	if err != nil {
		panic("Email not found, press enter to restart")
	}

	fmt.Println("\nPassword changed, press enter to back")
}
