package main

import (
	"log/slog"
	"net/http"

	unit "unit.nginx.org/go"
)

func main() {
	unit.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("hello world from dummyapp")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello, world from dummyapp"))
	}))
}