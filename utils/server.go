package utils

import "net/http"

func serve() {
	mux := http.NewServeMux()

	routing()

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
