package handlers

import (
	"Golang4/models"
	"bufio"
	"fmt"
	"strings"
)

func RegisterUser(service models.AuthService, scanner *bufio.Scanner) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("\nError: %v\n", r)
			waitEnter(scanner)
		}
	}()

	defer waitEnter(scanner)

	fmt.Println("\n---- Register ----")
	firstName := getInput("What is your first name: ", scanner)
	lastName := getInput("What is your last name: ", scanner)
	email := getInput("What is your email: ", scanner)
	password := getInput("Enter a strong password: ", scanner)
	confirmPassword := getInput("Confirm your password: ", scanner)

	if password != confirmPassword {
		panic("Password doesn't match!")
	}

	fmt.Println("\nIs it true?")
	fmt.Printf("First Name: %s\n", firstName)
	fmt.Printf("Last Name: %s\n", lastName)
	fmt.Printf("Email: %s\n", email)
	confirm := getInput("Continue (y/n): ", scanner)

	if strings.ToLower(confirm) != "y" {
		fmt.Println("\nRegistration cancelled.")
		return
	}

	err := service.Register(firstName, lastName, email, password)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("\nRegister success, press enter to back..")
}

func getInput(prompt string, scanner *bufio.Scanner) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func waitEnter(scanner *bufio.Scanner) {
	fmt.Print("Press enter to back ")
	scanner.Scan()
}
