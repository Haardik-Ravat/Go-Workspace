package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "whoa, it ran")

}

func abt_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "about")

}

func main() {

	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", abt_handler)
	http.ListenAndServe(":8000", nil)

}
