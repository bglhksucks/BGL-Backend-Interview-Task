package main

import (
	"log"
	"net/http"
	"os"

	"bgl/binance"
	"bgl/db"
	"bgl/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("APP_PORT")

	dbConn := db.Conn() // A single dbConn instance for all db connections
	defer dbConn.SQLDB.Close()

	dbConn.Init()
	go binance.GetBtcPriceEveryMinute(dbConn) // Start periodical backgound job

	router := server.InitRoutes(dbConn)
	http.ListenAndServe(":"+port, router)
}
