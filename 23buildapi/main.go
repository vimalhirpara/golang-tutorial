package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// /////////////////////// 1 Start ////////////////////////////

// Models
type Cource struct {
	CourceId    string  `json:"courceid"`
	CourceName  string  `json:"courcename"`
	CourcePrice int     `json:"-"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// DB
var cources []Cource

// Helper methods
func (c *Cource) IsEmpty() bool {
	return c.CourceName == ""
	//return c.CourceId == "" && c.CourceName == ""
}

///////////////////////// 1 End ////////////////////////////

func main() {
	fmt.Println("API - CRUD in slice")
	r := mux.NewRouter()

	// seeding
	cources = append(cources, Cource{CourceId: "2", CourceName: "React JS",
		CourcePrice: 299, Author: &Author{Fullname: "Vimal", Website: "xxx.xx"}})
	cources = append(cources, Cource{CourceId: "3", CourceName: "Angular JS",
		CourcePrice: 199, Author: &Author{Fullname: "Vimal", Website: "yyy.yy"}})

	// routing
	r.HandleFunc("/", serveHome).methods("GET")
	r.HandleFunc("/cources", getAllCources).methods("GET")
	r.HandleFunc("/cource/{id}", getOneCource).methods("GET")
	r.HandleFunc("/cource", createOneCource).methods("POST")
	r.HandleFunc("/cource/{id}", updateOneCource).methods("PUT")
	r.HandleFunc("/cource/{id}", deleteOneCource).methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

///////////////////////// 2 Start ////////////////////////////
// Controller - File
// server home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to api</h1>"))
}

func getAllCources(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all cources")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cources)
}

///////////////////////// 2 End ////////////////////////////

///////////////////////// 3 Start ////////////////////////////

func getOneCource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one cource")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop cources, find matching & send response

	for _, cource := range cources {
		if cource.CourceId == params["id"] {
			json.NewEncoder(w).Encode(cource)
			return
		}
	}

	json.NewEncoder(w).Encode("No cource found.")
	return
}

///////////////////////// 3 End ////////////////////////////

///////////////////////// 4 Start ////////////////////////////

func createOneCource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one cource.")
	w.Header().Set("Content-Type", "application/json")

	// if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data.")
		return
	}

	// if data is like {}
	var cource Cource
	_ = json.NewDecoder(r.Body).Decode(&cource)

	if cource.IsEmpty() {
		json.NewEncoder(w).Encode("no data inside json.")
		return
	}

	// generate unique id
	rand.Seed(time.Now().UnixNano())
	cource.CourceId = strconv.Itoa(rand.Intn(100)) // convert id then assign

	// add new cource in Cources
	cources = append(cources, cource)

	json.NewEncoder(w).Encode(cource)
	return
}

///////////////////////// 4 End ////////////////////////////

///////////////////////// 5 Start ////////////////////////////

func updateOneCource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one cource.")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop, get id, remove, add with id
	for index, cource := range cources {
		if cource.CourceId == params["id"] {
			// remove record
			cources = append(cources[:index], cources[index+1:]...)

			var cource Cource
			_ = json.NewDecoder(r.Body).Decode(&cource) // decode json
			cource.CourceId = params["id"]              // get id
			cources = append(cources, cource)
			json.NewEncoder(w).Encode(cource)
			return
		}
	}

	// when id is not found
}

///////////////////////// 5 End ////////////////////////////

///////////////////////// 6 Start ////////////////////////////

func deleteOneCource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one cource.")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop, get id, remove, add with id
	for index, cource := range cources {
		if cource.CourceId == params["id"] {
			// remove record
			cources = append(cources[:index], cources[index+1:]...)

			json.NewEncoder(w).Encode("Record is deleted with ", params["id"])
			return
			break
		}
	}
}

///////////////////////// 6 End ////////////////////////////
