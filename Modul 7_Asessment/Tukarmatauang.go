package main

import "fmt"

func main() {
	var qirat, dinar, dirham, fals int
	fmt.Scanln(&qirat)
	dinar = qirat/600
	dirham = qirat %600/60
	fals = qirat %600 % 60 /6
	qirat =qirat % 600 % 60 % 6
	fmt.Println(dinar,dirham,fals,qirat, " ")


}