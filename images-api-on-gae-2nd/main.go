package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/", indexHandler)

	log.Println("stating server...")
	if err := http.ListenAndServe(":8080", m); err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}
