package main

import "fmt"

func main() {
	name := "Zaki"

	switch name {
	case "Zaki":
		fmt.Println("Hello Zaki")
	case "Muhammad":
		fmt.Println("Hello Muhammad")
	default:
		fmt.Println("Hello Stranger")
	}

	// Example of short statement in switch
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Your name is quite long!")
	case false:
		fmt.Println("Your name is of average length.")
	}

	switch {
	case length > 10:
		fmt.Println("Your name is very long!")
	case length > 5:
		fmt.Println("Your name is quite long!")
	default:
		fmt.Println("Your name is of average length.")
	}
}