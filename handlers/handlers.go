package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

func LandingViewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	path := filepath.Join("..", "static", "html", "index.html")
	http.ServeFile(w, r, path)
	fmt.Println(vars)
}

func ForecastViewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	path := filepath.Join("..", "static", "html", "forecast.html")
	http.ServeFile(w, r, path)
	fmt.Println(vars)
}

func StateViewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	path := filepath.Join("..", "static", "html", "state.html")
	http.ServeFile(w, r, path)
	fmt.Println(vars)
}
