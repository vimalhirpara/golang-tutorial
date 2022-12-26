package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vimalhirpara/mongoapi/router"
)

func main() {
	fmt.Println("Mongodb Api")

	r := router.Router()

	fmt.Println("Server is getting started...")

	log.Fatal(http.ListenAndServe(":4000", r))

	fmt.Println("listening port 4000")
}
