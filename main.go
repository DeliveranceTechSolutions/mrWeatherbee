package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/deliveranceTechSolutions/mrWeatherbee/handlers"
)
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.LandingViewHandler).Methods("POST")
	r.HandleFunc("/forecast", handlers.ForecastViewHandler).Methods("POST")
	r.HandleFunc("/state", handlers.StateViewHandler).Methods("POST")
	http.Handle("/", r)
	
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Errorf("Error: ", err)
	}

	return
}
