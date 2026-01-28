package main

import "fmt"

func main() {
	// const firstName string = "Muhammad"
	// const lastName string = "Zaki"

	const (
		firstName string = "Muhammad"
		lastName  string = "Zaki"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(firstName + " " + lastName)

	//error
	// firstName = "As Shidiqi"
}