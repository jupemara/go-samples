package main

import (
	"log"
	"net/http"
)

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/", http.HandlerFunc(func(
		w http.ResponseWriter, r *http.Request,
	) {
		// for k, v := range r.Header {
		// 	for i, header := range v {
		// 		log.Printf(`%s[%d]: %s`, k, i, header)
		// 	}
		// }
		w.Write([]byte("content"))
	}))
	log.Fatal(http.ListenAndServe(":8080", m))
}
