package main

import "fmt"

func main() {

	type NoKTP string

	var ktpZaki NoKTP = "327682736872368723"

	var contoh string = "539125985198798714"
	var contohKTP NoKTP = NoKTP(contoh)

	fmt.Println(ktpZaki)
	fmt.Println(contohKTP)
}