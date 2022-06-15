package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		// handle err
	}

	// name := r.FormValue("test")

	fmt.Fprint(w, r.FormValue("test"))
	fmt.Println("endpoint hit: homepage")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome to the about page")
	fmt.Println("endpoint hit: aboutpage")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/about", aboutPage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	handleRequest()
}
