package main

import "fmt"

func main() {
	fmt.Println("Pointer")

	// var ptr *int
	// fmt.Println("Value of ptr is: ", ptr)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of actual ptr is:", ptr)
	fmt.Println("Value of actual ptr is:", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is", myNumber)
	fmt.Println("New value is", *ptr)
}
