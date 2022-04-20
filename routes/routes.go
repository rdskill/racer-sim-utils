package routes

import (
	"net/http"

	"github.com/rdskill/racer-sim-utils/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/calcular", controllers.Calcular)
}
