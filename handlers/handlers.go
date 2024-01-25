package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/deliveranceTechSolutions/mrWeatherbee/service/location"
)

func LandingViewHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	local := location.NewCore("WA")

	fmt.Println(local.Coordinates)
	path := filepath.Join("..", "static", "html", "index.html")
	http.ServeFile(w, r, path)
}

func ForecastViewHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	path := filepath.Join("..", "static", "html", "forecast.html")
	http.ServeFile(w, r, path)
}

func StateViewHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	path := filepath.Join("..", "static", "html", "state.html")
	http.ServeFile(w, r, path)
}
