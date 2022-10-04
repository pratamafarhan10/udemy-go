package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/marshal", marshal)
	http.HandleFunc("/encode", encode)
	http.ListenAndServe(":8080", nil)
}

type person struct {
	FirstName, LastName string
	Items               []string
}

func marshal(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		FirstName: "James",
		LastName:  "Bond",
		Items:     []string{"Guns", "Pistols", "Suit"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln(err)
	}

	w.Write(bs)
}

func encode(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		FirstName: "James",
		LastName:  "Bond",
		Items:     []string{"Guns", "Pistols", "Suit"},
	}

	// enc := json.NewEncoder(w)
	// err := enc.Encode(p1)
	err := json.NewEncoder(w).Encode(p1)

	if err != nil {
		log.Fatalln(err)
	}
}
