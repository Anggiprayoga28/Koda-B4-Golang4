package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Selesai")
	defer fmt.Println("Terima kasih")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("~~~~~SELAMAT DATANG~~~~~")

	for {
		fmt.Println("\n1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Lupa Password")
		fmt.Println("4. Keluar")
		fmt.Print("\nPilihan: ")

		scanner.Scan()
		pilihan := scanner.Text()

		switch pilihan {
		case "1":
			fmt.Println("Menu Register")
		case "2":
			fmt.Println("Menu Login")
		case "3":
			fmt.Println("Menu Lupa Password")
		case "4":
			fmt.Println("Keluar dari program")
			os.Exit(0)
		default:
			fmt.Println("Pilihan salah")
		}
	}
}
