package main

import "fmt"

func main() {
	var p1, p2, p3 string
	fmt.Scan(&p1, &p2, &p3)

	if p1 == p2 && p1 == p3 {
		fmt.Println("imbang")
	} else if p1 != p2 && p1 != p3 {
		fmt.Println("pemain 1 pemenang")
	} else if p2 != p1 && p2 != p3 {
		fmt.Println("pemain 2 pemenang")
	} else {
		fmt.Println("pemain 3 pemenang")
	}
}
