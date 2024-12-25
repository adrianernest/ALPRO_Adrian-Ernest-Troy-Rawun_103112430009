package main

import (
	"fmt"
	"math"
)

func main() {
	var k int

	fmt.Print("Nilai K = ")
	fmt.Scanln(&k)

	result := 1.0

	for i := 0; i < k; i++ {
		numerator := math.Pow(float64(4*i+2), 2)
		denominator := float64((4*i)*(4*i + 3))

		result *= numerator / denominator
	}

	fmt.Printf("Hasil dari operasi fungsi = %10f\n", result)
}
