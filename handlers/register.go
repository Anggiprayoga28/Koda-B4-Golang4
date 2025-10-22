package handlers

import (
	"Golang4/models"
	"bufio"
	"fmt"
	"strings"
)

func RegisterUser(users *[]*models.User, scanner *bufio.Scanner) {
	defer waitEnter(scanner)

	fmt.Println("\n---- Register ----")
	firstName := getInput("What is your first name: ", scanner)
	lastName := getInput("What is your last name: ", scanner)
	email := getInput("What is your email: ", scanner)
	password := getInput("Enter a strong password: ", scanner)
	confirmPassword := getInput("Confirm your password: ", scanner)

	if password != confirmPassword {
		fmt.Println("\nPassword doesn't match!")
		return
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

	for _, user := range *users {
		if user.Email == email {
			fmt.Println("\nEmail is used! Press enter to back..")
			return
		}
	}

	hashedPassword := models.HashPassword(password)

	newUser := &models.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  hashedPassword,
	}

	*users = append(*users, newUser)

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
