package main

import (
	"net/http"

	_ "github.com/lib/pq"
	"github.com/ryansweigart3/golang-rest-api/controller"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	getHandler := http.HandlerFunc(controller.GetPctResistYear)

	// Welcome path
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the API Server. \n /api/ar = all ar data \n /api/ar/year/{eventyearID} = all ar data by year"))
	})

	// Routes
	r.Route("/api", func(r chi.Router) {

		// AR Path
		r.Get("/ar", controller.GetPctResist)

		// AR year path
		r.Route("/ar/year", func(r chi.Router) {
			r.Use(controller.YearMiddleware)
			r.Get("/", getHandler)
		})

		// HAI Path
		r.Get("/hai", controller.GetPctResist)
	})
	http.ListenAndServe(":8080", r)
}
