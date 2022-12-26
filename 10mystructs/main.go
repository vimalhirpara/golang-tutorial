package main

import "fmt"

func main() {
	fmt.Println("welcome to my structs")

	vimal := User{"Vimal", "vimal@abc.com", true, 25}
	fmt.Println(vimal)
	fmt.Printf("Details are: %+v", vimal)
	fmt.Printf("\nName is %v and Email is %v", vimal.Name, vimal.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
