package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println("presentTime: ", presentTime)
	fmt.Println("Date: " + presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2022, time.November, 11, 06, 20, 0, 0, time.Local)
	fmt.Println("createdDate: ", createdDate)
	fmt.Println("Date: " + presentTime.Format("01-02-2006 Monday"))
}
