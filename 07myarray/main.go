package main

import "fmt"

func main() {
	fmt.Println("Welcome to array")

	var fruiltList [4]string

	fruiltList[0] = "Apple"
	fruiltList[1] = "Banana"
	fruiltList[3] = "Peach"

	fmt.Println("Fruit List is", fruiltList)
	fmt.Println("Total fruit", len(fruiltList))

	var vegList = [4]string{"Tomato", "Peace", "Beans"}

	fmt.Println("Veg List is", vegList)
	fmt.Println("Total veg", len(vegList))
}
