package main

import "fmt"

func main() {
	var name string

	name = "Muhammad"
	fmt.Println(name)

	name = "Zaki"
	fmt.Println(name)

	var lastName = "As Shidiqi"
	fmt.Println(lastName)

	firstName := "Muhammad"
	fmt.Println(firstName)

	var (
		first = "Muhammad"
		middle = "Zaki"
		last = "As Shidiqi"
	)
	fmt.Println(first)
	fmt.Println(middle)
	fmt.Println(last)
}