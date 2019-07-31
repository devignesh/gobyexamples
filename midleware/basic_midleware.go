package main

import (
	"fmt"
	"log"
	"net/http"
)

func vicky(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func smazz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "vicky_smazz")
}

func mca(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "vicky_mca")
}

func main() {
	http.HandleFunc("/smazz", vicky(smazz))
	http.HandleFunc("/mca", vicky(mca))

	http.ListenAndServe(":8080", nil)
}
