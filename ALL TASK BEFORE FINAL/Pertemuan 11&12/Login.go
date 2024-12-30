package main

import "fmt"

func main() {
	var username, password string
	failedAttempts := 0

	// Loop hingga login berhasil
	for {
		fmt.Print("Username: ")
		fmt.Scanln(&username)
		fmt.Print("Password: ")
		fmt.Scanln(&password)

		if username == "admin" && password == "admin" {
			break
		} else {
			failedAttempts++
		}
	}

	fmt.Printf("Jumlah kegagalan login: %d\n", failedAttempts)
	fmt.Println("Login berhasil")
}
