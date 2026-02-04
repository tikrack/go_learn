package main

import (
	"net/http"
)

func main() {
	server := http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, World!"))
		}),
	}

	println("Project Serving in http://localhost:8080")
	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
