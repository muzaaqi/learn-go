package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Muhammad"
	names[1] = "Zaki"
	names[2] = "As Shidiqi"

	fmt.Println(names[0]) // Muhammad
	fmt.Println(names[1]) // Zaki
	fmt.Println(names[2]) // As Shidiqi

	fmt.Println(len(names)) // 3

	var values = [...]int {
		90, 
		80, 
		70,
	}
	fmt.Println(values)    // [90 80 70]
	fmt.Println(values[0]) // 90
	fmt.Println(values[1]) // 80
	fmt.Println(values[2]) // 70

	fmt.Println(len(values)) // 3
}