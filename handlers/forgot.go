package handlers

import (
	"Golang4/models"
	"bufio"
	"fmt"
)

func ForgotPassword(service *models.UserService, scanner *bufio.Scanner) {
	defer waitEnter(scanner)

	fmt.Println("\n---- Forgot Password ----")
	fmt.Println()
	email := getInput("Enter your email: ", scanner)
	password := getInput("Enter a strong password: ", scanner)
	confirmPassword := getInput("Confirm your password: ", scanner)

	if password != confirmPassword {
		fmt.Println("\nPassword doesn't match!")
		return
	}

	err := service.ResetPassword(email, password)
	if err != nil {
		fmt.Println("\nEmail not found, press enter to restart")
		return
	}

	fmt.Println("\nPassword changed, press enter to back")
}
