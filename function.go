package main

import "fmt"

// normal function
func sayHello() {
	fmt.Println("Hello")
}

// function with parameter
func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

// function with return value
func getHello(name string) string {
	return "Hello " + name
}

// return multi value
func getFullName() (string, string) {
	return "Muhammad", "Zaki"
}

// named return value function
func getCompleteName() (firstName string, lastName string) {
	firstName = "Muhammad"
	lastName = "Zaki"
	return firstName, lastName
}

func main() {
	sayHello()
	sayHelloTo("Muhammad", "Zaki")
	result := getHello("Zaki")
	fmt.Println(result)
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
	firstName, lastName = getCompleteName()
	fmt.Println(firstName, lastName)
}