package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/company", func(res http.ResponseWriter, req *http.Request) {
		displayParameter(res, req)
	})
	println("Enter the url in browser:  http://localhost:8080/company?name=Tom&age=27")
	http.ListenAndServe(":8080", nil)
}
func displayParameter(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "name: "+req.FormValue("name"))
	io.WriteString(res, "\nage: "+req.FormValue("age"))
}
