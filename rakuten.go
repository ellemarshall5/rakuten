package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type Foo struct {
	name string `json:"name"`
	id   string `json:"id"`
}

type FooRecords []Foo

var records = FooRecords{
	{
		name: "test",
		id:   "test",
	},
}

func postFoo(w http.ResponseWriter, r *http.Request) {
	myuuid := uuid.NewV4()
	var newFoo Foo
	reqBody, _ := ioutil.ReadAll(r.Body)

	log.Print(myuuid)

	json.Unmarshal(reqBody, &newFoo)
	records = append(records, newFoo)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newFoo)
}

func getFoo(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["id"]

	for _, singleFoo := range records {
		if singleFoo.id == key {
			json.NewEncoder(w).Encode(singleFoo)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func deleteFoo(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["id"]

	for i, singleFoo := range records {
		if singleFoo.id == key {
			records = append(records[:i], records[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}

}

func main() {

	newRouter := mux.NewRouter().StrictSlash(true)
	newRouter.HandleFunc("/foo", postFoo).Methods("POST")
	newRouter.HandleFunc("/foo{id}", getFoo).Methods("GET")
	newRouter.HandleFunc("/foo{id}", deleteFoo).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", newRouter))

}
