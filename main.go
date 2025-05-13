package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.ServeFile(w, r, "./static/form.html")
		return 
	}

	err := r.ParseForm()

	if err != nil {
		fmt.Fprintf(w, "Form parse error: %v", err)
	}

	fmt.Fprintf(w, "POST request successful.\n")

	name := r.FormValue("name")
	email := r.FormValue("email")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name: %v\nEmail: %v\nAddress: %v\n", name, email, address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.URL.Path)

	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprint(w, "Hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Server is starting at port 8080\n")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
