package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	fmt.Println("asdasd")
	log.Println("Listening on localhost:8181")

	log.Fatal(http.ListenAndServe(":8181", nil))
}
