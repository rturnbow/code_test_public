package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	)

var tgt_port = 14000

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	
	if r.Method != "GET" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	
	fmt.Fprintf(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
	log.Println("Received request for "+name+" at "+address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)	
	http.HandleFunc("/form", formHandler)
	
	fmt.Println("Starting server at port "+strconv.Itoa(tgt_port)+" \n")
	
	if err := http.ListenAndServe(":"+strconv.Itoa(tgt_port), nil); err != nil {
		log.Fatal(err)
	}
}