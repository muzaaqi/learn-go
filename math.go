package main

import "fmt"

func main() {
	var a = 10
	var b = 20

	var c = a + b
	var d = a - b
	var e = a * b
	var f = a / b
	var g = a % b

	fmt.Println("a + b =", c)
	fmt.Println("a - b =", d)
	fmt.Println("a * b =", e)
	fmt.Println("a / b =", f)
	fmt.Println("a % b =", g)

	// augmented assignment
	var i = 10
	i += 10
	fmt.Println(i) // 20
	i -= 5
	fmt.Println(i) // 15
	i *= 2
	fmt.Println(i) // 30
	i /= 3
	fmt.Println(i) // 10
	i %= 7
	fmt.Println(i) // 3

	// unary operator
	var j = 10
	j++
	fmt.Println(j) // 11
	j--
	fmt.Println(j) // 10
	
}