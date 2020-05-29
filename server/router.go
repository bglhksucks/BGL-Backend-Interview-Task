package server

import (
	"bgl/db"

	"github.com/go-chi/chi"
)

var dbConn *db.DB

// InitRoutes hello
func InitRoutes(db *db.DB) *chi.Mux {
	dbConn = db
	r := chi.NewRouter()

	r.Get("/lastprice", getLastPrice)
	r.Get("/price/{time}", getPriceAtTime)
	r.Get("/averagePrice/{from}/{to}", getAveragePrice)

	return r
}
