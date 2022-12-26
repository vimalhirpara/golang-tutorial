package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input program"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin) // init reader
	fmt.Println("Enter the rating for pizza:")

	// comma ok || err err
	input, _ := reader.ReadString('\n') // listen reader
	fmt.Println("Thanks for rating", input)
	fmt.Printf("Variable type is: %T", input)
}
