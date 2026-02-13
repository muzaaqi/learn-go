package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	goodbye := getGoodBye
	result := goodbye("Zaki")
	fmt.Println(result)
}