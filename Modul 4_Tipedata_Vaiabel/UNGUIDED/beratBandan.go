package main 

import "fmt"

func main () {
	var TB, BB, BMI float64

	fmt.Print("Masukan BMI nya: ")
	fmt.Scan(&BMI)

	fmt.Print("Masukan tinggi badan:")
	fmt.Scan(&TB)

	BB = BMI *  (TB * TB)

	fmt.Print("Jadi berat badannya adalah:", BB)

}