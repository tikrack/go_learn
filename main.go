package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<h1>Hello</h1`))
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	println("Project Serving in http://localhost:8080")
	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
