package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}

	// person["name"] = "Muhammad Zaki As Shidiqi"
	// person["address"] = "Indonesia"
	// person["job"] = "Programmer"


	person := map[string]string{
		"name":    "Muhammad Zaki As Shidiqi",
		"address": "Indonesia",
		"job":     "Programmer",
	}

	fmt.Println(person)            // map[address:Indonesia job:Programmer name:Muhammad Zaki As Shidiqi]
	fmt.Println(person["name"])    // Muhammad Zaki As Shidiqi
	fmt.Println(person["address"]) // Indonesia
	fmt.Println(person["job"])     // Programmer
	fmt.Println(len(person))       // 3

	book := make(map[string]string)
	book["title"] = "Learn Go Programming"
	book["author"] = "John Doe"
	book["ups"] = "Salah"

	fmt.Println(book)            // map[author:John Doe title:Learn Go Programming ups:Salah]
	fmt.Println(book["title"])		// Learn Go Programming
	fmt.Println(book["author"])		// John Doe
	fmt.Println(len(book))       // 3
	delete(book, "ups")
	fmt.Println(book)            // map[author:John Doe title:Learn Go Programming]
}