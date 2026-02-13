package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Counter:", counter)
	// 	counter++
	// }
	
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Counter:", counter)
	}

	fmt.Println("Finished first loop")

	names := []string{"Zaki", "Muhammad", "Ali"}
	// manual
	for i := 0; i < len(names); i++ {
		fmt.Println("Name", i, "=", names[i])
	}

	// range
	for index, name := range names {
		fmt.Println("Name", index, "=", name)
	}
	
	for _, name := range names {
		fmt.Println(name)
	}
}