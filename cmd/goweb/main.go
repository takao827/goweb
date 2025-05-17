package main

import (
	"fmt"
	"net/http"

	unit "unit.nginx.org/go"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World, %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	unit.ListenAndServe(":8080", nil)
}
