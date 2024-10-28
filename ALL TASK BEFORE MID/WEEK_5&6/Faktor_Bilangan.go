package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Printf("%d %t\n", i, true)
		} else {
			fmt.Printf("%d %t\n", i, false)
		}
	}
}