package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/deliveranceTechSolutions/mrWeatherbee/handlers"
)
func main() {
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.HandleFunc("/", handlers.LandingViewHandler)
	r.HandleFunc("/forecast", handlers.ForecastViewHandler)
	r.HandleFunc("/state", handlers.StateViewHandler)

	http.Handle("/", r)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Errorf("Error: ", err)
	}

	return
}
