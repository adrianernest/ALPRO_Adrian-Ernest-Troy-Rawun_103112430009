package main

import "fmt"

func main () {
	var x, j int
	fmt.Scanln(&x)
	for j=1 ; j<=x ; j+=1 {
		if x  %j == 0 {
			fmt.Print(j," ")
		}
	}
}