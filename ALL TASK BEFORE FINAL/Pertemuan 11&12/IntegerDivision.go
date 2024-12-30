package main

import "fmt"

func main () {
	var x, xi, y, i int
	fmt.Scan(&x,&y)
	i = 0
	xi = for xi >=y  {
		xi = xi -y
		i = i + xi
	}
	fmt.Println(i)
}