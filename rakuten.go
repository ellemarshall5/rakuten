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
	Name string `json:"name"`
	Id   string `json:"id"`
}

var records []Foo

func postFoo(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var foo Foo
	json.Unmarshal(reqBody, &foo)
	foo.Id = uuid.NewV4().String()
	records = append(records, foo)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foo)

}

func getFoo(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["id"]
	var found bool

	if len(records) == 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		for _, singleFoo := range records {
			if singleFoo.Id == key {
				found = true
				json.NewEncoder(w).Encode(singleFoo)
				w.WriteHeader(http.StatusOK)
				break
			}
		}
		if !found {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func deleteFoo(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["id"]
	var found bool

	if len(records) == 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		for i, singleFoo := range records {
			if singleFoo.Id == key {
				found = true
				records = append(records[:i], records[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
				break
			}
		}
		if !found {
			w.WriteHeader(http.StatusNotFound)
		}
	}

}

func main() {
	newRouter := mux.NewRouter().StrictSlash(true)
	newRouter.HandleFunc("/foo", postFoo).Methods("POST")
	newRouter.HandleFunc("/foo/{id}", getFoo).Methods("GET")
	newRouter.HandleFunc("/foo/{id}", deleteFoo).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", newRouter))

}
