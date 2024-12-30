package main

import "fmt"

func main() {
	var val int
	fmt.Print("Masukan Nilai: ")
	fmt.Scanln(&val)

	if val > 0 {
		fmt.Print("Number is Positive")
		if val%2 == 0 {
			fmt.Print(" and Even")
		} else {
			fmt.Print(" and Odd")
		}
	} else if val < 0 {
		fmt.Print("Number is Negative")
	} else {
		fmt.Print("Number is Zero")
	}
	fmt.Print("\nFinish")
}