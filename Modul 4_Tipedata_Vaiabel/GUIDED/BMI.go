package main

import "fmt"

func main() {
	var TB, BB, BMI float64

	fmt.Print("masukan BB:")
	fmt.Scan(&BB)
	fmt.Print("masukan TB:")
	fmt.Scan(&TB)
	BMI = BB / (TB * TB)
	fmt.Printf("%.2f", BMI)
}
