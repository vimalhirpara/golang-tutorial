package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slice")

	var frtList = []string{"Apple", "peach", "Beans"}

	fmt.Printf("Type of list is %T\n", frtList)
	fmt.Println(frtList)

	frtList = append(frtList, "mango", "tomato")
	fmt.Println(frtList)

	frtList = append(frtList[1:3])
	fmt.Println(frtList)

	highScores := make([]int, 4)
	highScores[0] = 123
	highScores[1] = 234
	highScores[2] = 345
	highScores[3] = 456
	//highScores[0] = 567

	highScores = append(highScores, 678, 789, 567)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	var courses = []string{"react", "js", "swift", "python", "ruby"}
	fmt.Println("\n", courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)
}
