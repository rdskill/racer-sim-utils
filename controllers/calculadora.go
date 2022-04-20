package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/rdskill/racer-sim-utils/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", nil)
}

func Calcular(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		tempoVolta := r.FormValue("tempoVolta")
		duracao := r.FormValue("duracao")
		consumoVolta := r.FormValue("consumoVolta")
		maisUmaVolta := contains(r.Form["maisUmaVoltaCheck"], "maisUmaVolta")

		if tempoVolta == "" || duracao == "" || consumoVolta == "" {
			log.Println("Todos os campos devem ser preenchidos.")
			mostrarErro(w, "Todos os campos devem ser preenchidos.")
			return
		}

		if strings.Contains(consumoVolta, ",") {
			consumoVolta = strings.Replace(consumoVolta, ",", ".", -1)
		}

		consumoVoltaConvertidoParaFloat, err := strconv.ParseFloat(consumoVolta, 64)
		if err != nil {
			log.Println("Erro na conversão do consumo:", err)
			mostrarErro(w, "Erro na conversão do consumo. Valor: "+consumoVolta)
			return
		}

		duracaoConvertidaParaInt, err := strconv.Atoi(duracao)
		if err != nil {
			log.Println("Erro na conversão da duração:", err)
			mostrarErro(w, "Erro na conversão da duração. Valor: "+duracao)
			return
		}

		if !strings.Contains(tempoVolta, ":") {
			tempoVolta = tempoVolta + ":00"
		}

		tempoSplit := strings.Split(tempoVolta, ":")

		minutosVoltaEmSegundos, err := strconv.Atoi(tempoSplit[0])
		if err != nil {
			log.Println("Erro na conversão de minutos:", err)
			mostrarErro(w, "Erro na conversão de minutos. Valor: "+tempoSplit[0])
			return
		}

		segundosVolta, err := strconv.Atoi(tempoSplit[1])
		if err != nil {
			log.Println("Erro na conversão de segundos:", err)
			mostrarErro(w, "Erro na conversão de segundos. Valor: "+tempoSplit[1])
			return
		}

		tempoVoltaSegundos := (minutosVoltaEmSegundos * 60) + segundosVolta
		duracaoSegundos := duracaoConvertidaParaInt * 60
		voltas := duracaoSegundos / tempoVoltaSegundos

		if maisUmaVolta {
			voltas = voltas + 1
		}

		combustivelNecessario := float64(voltas) * consumoVoltaConvertidoParaFloat

		maisUmaVoltaStr := "Sim"

		if !maisUmaVolta {
			maisUmaVoltaStr = "Não"
		}

		calculadora := models.Calculadora{
			TempoVolta:            tempoVolta,
			Duracao:               duracaoConvertidaParaInt,
			ConsumoVolta:          consumoVoltaConvertidoParaFloat,
			Voltas:                voltas,
			CombustivelNecessario: combustivelNecessario,
			MaisUmaVolta:          maisUmaVoltaStr,
		}

		temp.ExecuteTemplate(w, "CalculadoraResult", calculadora)
	}
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	_, ok := set[item]
	return ok
}

func criarErro(mensagem, href string) models.Erro {
	return models.Erro{
		Mensagem: mensagem,
		Href:     href,
	}
}

func mostrarErro(w http.ResponseWriter, mensagem string) {
	temp.ExecuteTemplate(w, "Error", criarErro(mensagem, "/"))
}
