package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	result := sumAll(10, 20, 30, 40, 50)
	fmt.Println("Total:", result)
	result = sumAll(1, 2, 3)
	fmt.Println("Total:", result)


	numbers := []int{5, 10, 15, 20}
	result = sumAll(numbers...)
	fmt.Println("Total:", result)
}