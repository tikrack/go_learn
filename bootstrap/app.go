package bootstrap

import (
	"main/routes"
	"net/http"

	"github.com/fatih/color"
)

func Run() {
	mux := http.NewServeMux()

	routes.Api(mux)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	color.Green("\nStarting server to http://localhost:8080\n")
	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
