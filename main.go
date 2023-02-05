package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("http://127.0.0.1:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe error", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)

	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "406 Not Acceptable", http.StatusNotAcceptable)
	}
	fmt.Fprint(w, "Hello, world!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parser Error %v", err)
		log.Fatal("ParseForm", err)
	}
	name := r.Form["name"][0]
	address := r.Form["address"][0]

	fmt.Fprintf(w, "Address: %v, Name : %v ", address, name)
}
