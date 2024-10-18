package main

import "fmt"

func main() {
	var alas, tinggi, n, i int
	var luas float64
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&alas, &tinggi)
		luas = 0.5 * float64(alas) * float64(tinggi)
		fmt.Println(luas)
	}
}
