package main

import "fmt"

func main() {
	var kar string
	fmt.Println("Masukan Karakter:")

	fmt.Scan(&kar)
	if kar <= "a" && kar >= "z" || kar <= "A" && kar >= "Z" {
		fmt.Println("karakter ini  termasuk alphabet")
	} else {
		fmt.Println("karakter ini bukan termasuk alphabet")
	}

}
