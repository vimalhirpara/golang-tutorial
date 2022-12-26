package main

import "fmt"

const LoginToken string = "qwerty"

func main() {
	var username string = "vimal"
	fmt.Println(username)
	fmt.Printf("Variable type is: %T", username)

	fmt.Println("\n")

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable type is: %T", isLoggedIn)

	fmt.Println("\n")

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable type is: %T", smallVal)

	fmt.Println("\n")

	var smallFloat float64 = 256.12345678
	fmt.Println(smallFloat)
	fmt.Printf("Variable type is: %T", smallFloat)

	fmt.Println("\n")

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable type is: %T", anotherVariable)

	fmt.Println("\n")

	//implicit type
	var website = "google.com"
	fmt.Println(website)
	fmt.Printf("Variable type is: %T", website)

	fmt.Println("\n")

	// no var style
	numberOfUser := 30000
	fmt.Println(numberOfUser)

	fmt.Println("Login Token: " + LoginToken)
	fmt.Printf("Variable type is: %T", LoginToken)
}
