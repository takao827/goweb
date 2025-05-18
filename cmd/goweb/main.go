package main

import (
	"net/http"

	unit "unit.nginx.org/go"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)

	s := &http.Server{
		Addr:    config.Address,
		Handler: mux,
	}

	unit.ListenAndServe(s.Addr, s.Handler)
}
