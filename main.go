package main

import (
	"DomaiNesia-Pretest/app"
	"encoding/json"
	"net/http"
	"strconv"
)

/*array/slice for hold dummy data
with format Dvd model (in this case i use struct) */
var Dvds = []app.Dvd{}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		_ = json.NewEncoder(w).Encode(Dvds)
	} else {
		w.WriteHeader(405)
		_ = json.NewEncoder(w).Encode("Method Not Allowed")
	}
}

func Show(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//get items
	items := []app.Dvd{}

	//get params
	Ids, _ := r.URL.Query()["id"]
	Id := Ids[0]                 //get only in first array
	getId, _ := strconv.Atoi(Id) //convert to int

	for _, item := range Dvds {
		if item.Id == getId {
			items = append(items, item)
			break
		} else {
			items = nil
		}
	}

	if r.Method == "GET" {
		if items == nil {
			w.WriteHeader(404)
			_ = json.NewEncoder(w).Encode("The item is empty")
		} else {
			_ = json.NewEncoder(w).Encode(items)
		}
	} else {
		w.WriteHeader(405)
		_ = json.NewEncoder(w).Encode("Method Not Allowed")
	}
}

func main() {
	//dummy data for dvd
	Dvds = append(Dvds, app.Dvd{
		Id:       1,
		Title:    "Spongebob Squarepants",
		Price:    10000,
		Quantity: 10,
		Category: "cartoon",
	})
	Dvds = append(Dvds, app.Dvd{
		Id:       2,
		Title:    "Conjuring",
		Price:    25000,
		Quantity: 5,
		Category: "horror",
	})
	Dvds = append(Dvds, app.Dvd{
		Id:       3,
		Title:    "Interstellar",
		Price:    30000,
		Quantity: 15,
		Category: "fiction",
	})
	Dvds = append(Dvds, app.Dvd{
		Id:       4,
		Title:    "The Avengers",
		Price:    50000,
		Quantity: 5,
		Category: "action/fiction",
	})
	Dvds = append(Dvds, app.Dvd{
		Id:       5,
		Title:    "Three Idiots",
		Price:    20000,
		Quantity: 10,
		Category: "comedy",
	})

	http.HandleFunc("/dvds", Index)
	http.HandleFunc("/dvd", Show)

	//serve a server
	_ = http.ListenAndServe(":8000", nil)
}
