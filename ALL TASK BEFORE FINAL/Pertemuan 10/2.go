package main

import "fmt"

func main() {
	var a float32
	fmt.Print("Masukan TAK: ")
	fmt.Scan(&a)

	if a >= 0 {
		fmt.Print("Bilangan Positiv\n")
		if a < 2 {
			fmt.Print("Poor")
		} else if a >= 2 && a <= 2.75 {
			fmt.Print("Fair")
		} else if a >= 2.76 && a <= 3.00 {
			fmt.Print("Sastifactoy")
		} else if a >= 3.01 && a <= 3.50 {
			fmt.Print("Very Good")
		} else if a >= 3.50 {
			fmt.Print("Excellent")
		}
	} else {
		fmt.Print("Nomor salah")
	}
}