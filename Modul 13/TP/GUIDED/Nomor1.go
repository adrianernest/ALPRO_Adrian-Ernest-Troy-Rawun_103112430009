package main

import "fmt"

func main() {	
	var bilangan int

	fmt.Print("Masukan bilangan : ")
	fmt.Scan(&bilangan)

	for i := 1; i <= bilangan; i++ {
		if i%2 !=0 {
			fmt.Println(i)
		}
	}
}