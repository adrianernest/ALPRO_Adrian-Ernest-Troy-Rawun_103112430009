package main

import (
	"fmt"
	"math"
)

func main() {
	var fx float64

	fmt.Print("Masukkan nilai f(x): ")
	fmt.Scanln(&fx)

	// Hitung nilai x dari persamaan f(x) = 2/(x+5) + 5
	x := (2 / (fx - 5)) - 5

	// Cek jika nilai x valid (tidak tak terhingga)
	if !math.IsInf(x, 0) {
		fmt.Println("Nilai x:", x)
	} else {
		fmt.Println("Nilai x tidak valid.")
	}
}