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
	param1 := r.URL.Query().Get("eventyearid")
	param1s := r.URL.Query()["param1"]
	// param := r.URL.Query().Get("eventyearid")
	// param2 := r.URL.Query()["state"]
	// eventyearid := chi.URLParam(r, "eventyearid")

	rows, err := db.Query("SELECT * FROM antibiotic_resistance")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var arCube []models.AntibioticResistance

	for rows.Next() {
		var artext models.AntibioticResistance
		rows.Scan(&artext.EventYearID, &artext.State, &artext.Pct_Resist)
		if param1 == artext.EventYearID && len(param1s) > 0 {
			arCube = append(arCube, artext)
		}
		// if param1 == artext.EventYearID {
		// 	arCube = append(arCube, artext)
		// } else if param2 == artext.State {
		// 	arCube = append(arCube, artext)
		// }
	}

	arBytes, _ := json.MarshalIndent(arCube, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(arBytes)

	defer db.Close()
}
