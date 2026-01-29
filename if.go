package main

import "fmt"

func main() {
	name := "Zaki"

	if name == "Zaki" {
		fmt.Println("Hello Zaki")
	} else if name == "Muhammad" {
		fmt.Println("Hello Muhammad")
	} else {
		fmt.Println("Hello Stranger")
	}

	// Example of short statement in if
	if length := len(name); length > 5 {
		fmt.Println("Your name is quite long!")
	} else {
		fmt.Println("Your name is of average length.")
	}
}