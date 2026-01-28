package main

import "fmt"

func main() {
	names := [...]string{"Alice", "Bob", "Charlie", "Diana", "Eve", "Frank", "Grace", "Heidi", "Ivan", "Judy"}

	slice1 := names[4:6]
	fmt.Println(slice1)
	fmt.Println(len(slice1)) // Length of the slice
	fmt.Println(cap(slice1)) // Capacity of the slice

	slice2 := names[:3]
	fmt.Println(slice2)
	fmt.Println(len(slice2)) // Length of the slice
	fmt.Println(cap(slice2)) // Capacity of the slice

	slice3 := names[3:]
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // Length of the slice
	fmt.Println(cap(slice3)) // Capacity of the slice

	slice4 := names[:]
	fmt.Println(slice4)
	fmt.Println(len(slice4)) // Length of the slice
	fmt.Println(cap(slice4)) // Capacity of the slice

	// appending to a slice
	days := []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:] 
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	// daysBaru := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu Baru", "Minggu Baru", "Libur Baru"}
	daySlice2[0] = "Sabtu Lama"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	// creating a new slice with make
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Hello"
	newSlice[1] = "World"
	// newSlice[2] = "Golang" --> This will cause a runtime panic: index out of range
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice = append(newSlice, "Golang")
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice[0] = "Hi"
	fmt.Println(newSlice)

	// copying slices
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy (toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// declaring array and slice
	thisArray := [...]string{"A", "B", "C", "D", "E"}
	thisSlice := []string{"X", "Y", "Z"}

	fmt.Println(thisArray)
	fmt.Println(thisSlice)
}