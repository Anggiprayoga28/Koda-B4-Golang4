package main

import (
	"Golang4/handlers"
	"Golang4/models"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scanner *bufio.Scanner

func main() {
	defer fmt.Println("\nThank you for using our system!")

	scanner = bufio.NewScanner(os.Stdin)
	var users []*models.User

	for {
		showMainMenu()
		choice := getInput("Choose a menu: ")

		switch choice {
		case "1":
			handlers.RegisterUser(&users, scanner)
		case "2":
			fmt.Println("Menu Login")
		case "3":
			fmt.Println("Menu Lupa Password")
		case "0":
			fmt.Println("\nGoodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice!")
		}
	}
}

func showMainMenu() {
	defer fmt.Println()

	fmt.Println("\n---- Welcome to system ----")
	fmt.Println("\n1. Register")
	fmt.Println("2. Login")
	fmt.Println("3. Forgot Password")
	fmt.Println("\n0. Exit")
}

func getInput(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}
