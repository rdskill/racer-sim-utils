package main

import (
	"net/http"
	"os"

	"github.com/rdskill/racer-sim-utils/routes"
)

func main() {
	routes.CarregaRotas()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	http.ListenAndServe(":"+port, nil)
}
