package main

import (
	"log"
	"net/http"
	"roadmaps/projects/unit-converter/cmd"
)

func main() {
	http.HandleFunc("/convert", cmd.ConvertHandler)
	http.HandleFunc("/", cmd.HomePage)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
