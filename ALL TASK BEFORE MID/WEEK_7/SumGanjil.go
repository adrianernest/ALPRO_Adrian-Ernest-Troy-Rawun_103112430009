package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Masukkan a dan b: ")
	fmt.Scanln(&a, &b)

	sum := 0
	for i := a; i <= b; i++ {
		if i%2 != 0 {
			sum += i
		}
	}

	fmt.Println("Keluaran:", sum)
}
