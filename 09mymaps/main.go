package main

import "fmt"

func main() {
	fmt.Println("welcome to my maps")

	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["RB"] = "ruby"
	languages["PY"] = "python"

	fmt.Println(languages)
	fmt.Println(languages["JS"])

	delete(languages, "RB")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Printf("For key %v is %v\n", key, value)
	}

	for _, value := range languages {
		fmt.Printf("For key v is %v\n", value)
	}
}
