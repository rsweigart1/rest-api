package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"github.com/ryansweigart3/golang-rest-api/db"
	"github.com/ryansweigart3/golang-rest-api/models"
)

func GetPctResist(w http.ResponseWriter, r *http.Request) {
	db := db.OpenConnection()

	rows, err := db.Query("SELECT * FROM antibiotic_resistance")
	if err != nil {
		log.Fatal(err)
	}

	var arCube []models.AntibioticResistance

	for rows.Next() {
		var artext models.AntibioticResistance
		rows.Scan(&artext.EventYearID, &artext.State, &artext.Pct_Resist)
		arCube = append(arCube, artext)
	}

	arBytes, _ := json.MarshalIndent(arCube, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(arBytes)

	defer rows.Close()
	defer db.Close()
}

// Get year middleware
func YearMiddleware(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		chi.URLParam(r, "eventyearid")
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func GetPctResistYear(w http.ResponseWriter, r *http.Request) {
	db := db.OpenConnection()
	query := r.URL.Query()
	param := query.Get("eventyearid")
	// eventyearid := chi.URLParam(r, "eventyearid")

	rows, err := db.Query("SELECT * FROM antibiotic_resistance WHERE eventyearid=$1", param)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var arCube []models.AntibioticResistance

	for rows.Next() {
		var artext models.AntibioticResistance
		rows.Scan(&artext.EventYearID, &artext.State, &artext.Pct_Resist)
		// if eventyearid == artext.EventYearID {
		arCube = append(arCube, artext)
	}

	arBytes, _ := json.MarshalIndent(arCube, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(arBytes)

	defer db.Close()
}
