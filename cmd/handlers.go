package cmd

import (
	"html/template"
	"log"
	"net/http"
	"roadmaps/projects/unit-converter/internal/converter"
	"strconv"
)

var tpl = template.Must(template.ParseFiles("web/templates/layout.html"))

func HomePage(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func ConvertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fromUnit := r.FormValue("from")
		toUnit := r.FormValue("to")
		value, _ := strconv.ParseFloat(r.FormValue("value"), 64)
		var result float64

		switch r.FormValue("type") {
		case "length":
			result = converter.ConvertLength(value, fromUnit, toUnit)
		case "weight":
			result = converter.ConvertWeight(value, fromUnit, toUnit)
		case "temperature":
			result = converter.ConvertTemperature(value, fromUnit, toUnit)
		}

		tpl.Execute(w, result)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
