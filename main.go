package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "URL Not Matched", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is Invalid", http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "Hello Ketan")

}

func handleForm(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w, "Error Occured in Parsing Form", http.StatusBadRequest)
		return

	}

	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "name %v", name)
	fmt.Fprintf(w, "address %v", address)

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", handleForm)
	http.HandleFunc("/hello", handleHello)
	fmt.Println("Server Started at 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
