package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {
	value := User{Firstname: "vicky", Lastname: "kumar", Age: 23 }
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		
		json.NewDecoder(r.Body).Decode(&value)

		fmt.Fprintf(w, "%s %s is %d years old!", value.Firstname, value.Lastname, value.Age)
	})

	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		peter := User{
			Firstname: "vignesh",
			Lastname:  "kumar",
			Age:       23,
		}

		json.NewEncoder(w).Encode(peter)
	})

	http.ListenAndServe(":8080", nil)
}
