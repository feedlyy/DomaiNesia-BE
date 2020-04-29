package main

import (
	"DomaiNesia-BE/app"
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
			_ = json.NewEncoder(w).Encode("Wrong ID")
		} else {
			_ = json.NewEncoder(w).Encode(items)
		}
	} else {
		w.WriteHeader(405)
		_ = json.NewEncoder(w).Encode("Method Not Allowed")
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//get client/user input
	title := r.FormValue("title")
	price, _ := strconv.Atoi(r.FormValue("price"))       //convert input from string to int
	quantity, _ := strconv.Atoi(r.FormValue("quantity")) //convert input from string to int
	category := r.FormValue("category")

	//getLastId from slice
	lastId := Dvds[len(Dvds)-1].Id

	//slice to display in response
	slice := app.Dvd{
		Id:       lastId + 1, // always increment from the last id in slice Dvd's
		Title:    title,
		Price:    price,
		Quantity: quantity,
		Category: category,
	}

	//added to array dvd
	Dvds = append(Dvds, slice)

	if r.Method == "POST" {
		_ = json.NewEncoder(w).Encode(slice)
	} else {
		w.WriteHeader(405)
		_ = json.NewEncoder(w).Encode("Method Not Allowed")
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//get items
	items := []app.Dvd{}

	//get params
	Ids, _ := r.URL.Query()["id"]
	Id := Ids[0]                 //get only in first array
	getId, _ := strconv.Atoi(Id) //convert to int

	//get client/user input
	title := r.FormValue("title")
	price := r.FormValue("price")       //convert input from string to int
	quantity := r.FormValue("quantity") //convert input from string to int
	category := r.FormValue("category")

	for i := 0; i < len(Dvds); i++ {
		if Dvds[i].Id == getId {
			if title != "" {
				Dvds[i].Title = title
			}
			if price != "" {
				getPriceUpdate, _ := strconv.Atoi(r.FormValue("price"))
				Dvds[i].Price = getPriceUpdate
			}
			if quantity != "" {
				getQuantityUpdate, _ := strconv.Atoi(r.FormValue("quantity"))
				Dvds[i].Quantity = getQuantityUpdate
			}
			if category != "" {
				Dvds[i].Category = category
			}
			items = append(items, Dvds[i])
			break
		} else {
			items = nil
		}
	}

	if r.Method == "PUT" {
		if items == nil {
			w.WriteHeader(404)
			_ = json.NewEncoder(w).Encode("Wrong ID")
		} else {
			_ = json.NewEncoder(w).Encode(items)
		}
	} else {
		w.WriteHeader(405)
		_ = json.NewEncoder(w).Encode("Method Not Allowed")
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//get params
	Ids, _ := r.URL.Query()["id"]
	Id := Ids[0]                 //get only in first array
	getId, _ := strconv.Atoi(Id) //convert to int

	//get deleted item first
	items := []app.Dvd{}

	for i := 0; i < len(Dvds); i++ {
		if Dvds[i].Id == getId {
			items = append(items, Dvds[i])
			Dvds = append(Dvds[:i], Dvds[i+1:]...) //delete the item
			break
		} else {
			items = nil
		}
	}

	if r.Method == "DELETE" {
		if items == nil {
			w.WriteHeader(404)
			_ = json.NewEncoder(w).Encode("Wrong ID")
		} else {
			_ = json.NewEncoder(w).Encode("Success delete item")
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

	//list of routes
	/*please take a note for show,update,delete
	they need params in their url, so please do this when using one of them
	example: localhost:8000/show?id=1*/
	http.HandleFunc("/index", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)

	//serve a server
	_ = http.ListenAndServe(":8000", nil)
}
