package bootstrap

import (
	"main/routes"
	"net/http"
)

func Run() {
	mux := http.NewServeMux()

	routes.Api(mux)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
