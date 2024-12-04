package main

import 
	"fmt"


func main() {
	var word string
	var repetition int
	fmt.Scan(&word, &repetition)
	counter := 0
	for done := false; !done; {
	fmt.Println(word)
	counter++
	done = (counter >= repetition)
	}
}