package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	// Perulangan untuk terus meminta input hingga pengguna mengetik "telkom"
	for {
		// Meminta input dari pengguna
		fmt.Print("Masukkan kata: ")
		fmt.Scanln(&input)

		// Memeriksa apakah input adalah "telkom", tanpa peka huruf besar/kecil
		if strings.ToLower(input) == "telkom" {
			fmt.Println("Program selesai.")
			break // Keluar dari perulangan jika input adalah "telkom"
		} else {
			// Menampilkan input yang dimasukkan pengguna
			fmt.Printf("Anda mengetik: %s\n", input)
		}
	}
}
