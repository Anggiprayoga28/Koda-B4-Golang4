package handlers

import (
	"Golang4/models"
	"bufio"
	"fmt"
	"os"
)

func LoginUser(service models.AuthService, scanner *bufio.Scanner) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("\nError: %v\n", r)
			waitEnter(scanner)
		}
	}()

	defer waitEnter(scanner)

	fmt.Println("\n---- Login ----")
	fmt.Println()
	email := getInput("Enter your email: ", scanner)
	password := getInput("Enter your password: ", scanner)

	user, err := service.Login(email, password)
	if err != nil {
		panic("Wrong email or password, press enter to restart")
	}

	fmt.Println("\nLogin success, press enter to back..")
	waitEnter(scanner)

	showUserMenu(user, service, scanner)
}

func showUserMenu(user models.UserInterface, service models.AuthService, scanner *bufio.Scanner) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("\nError: %v\n", r)
		}
	}()

	for {
		fmt.Println("\n---- Welcome to system ----")
		fmt.Printf("\nHello %s!\n", user.GetEmail())
		fmt.Println("1. List All Users")
		fmt.Println("2. Logout")
		fmt.Println("\n0. Exit")
		fmt.Println()

		choice := getInput("Choose a menu: ", scanner)

		switch choice {
		case "1":
			listAllUsers(service, scanner)
		case "2":
			fmt.Println("\nLogged out successfully!")
			return
		case "0":
			fmt.Println("\nGoodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice!")
		}
	}
}

func listAllUsers(service models.AuthService, scanner *bufio.Scanner) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("\nError: %v\n", r)
		}
	}()

	defer waitEnter(scanner)

	fmt.Println("\n---- List all users ----")

	users := service.GetAllUsers()

	if len(users) == 0 {
		fmt.Println("No users registered yet.")
	} else {
		for i, user := range users {
			fmt.Printf("%d.\n", i+1)
			fmt.Printf("Full Name: %s\n", user.GetFullName())
			fmt.Printf("Email: %s\n", user.GetEmail())
			fmt.Printf("Password: %s\n\n", user.GetPassword())
		}
	}
}
