package db

import (
	"bgl/helpers"
	"database/sql"
	"fmt"
	"os"

	// For mysql connection
	_ "github.com/go-sql-driver/mysql"
)

// DB for database connection
type DB struct {
	SQLDB *sql.DB
}

// Conn intializes a DB connection
func Conn() *DB {
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_DATABASE")
	dbConnect := os.Getenv("DB_CONNECT")

	dbConn, err := sql.Open("mysql", dbUser+":"+dbPassword+"@tcp("+dbConnect+")/"+dbDatabase)
	//dbConn, err := sql.Open("mysql", dbUser+":"+dbPassword+"@tcp(mariadb)/"+dbDatabase)
	if err != nil {
		fmt.Println(err.Error())
	}

	db := &DB{dbConn}
	return db
}

// Init btc_price table if the table does not exist yet
func (db *DB) Init() {
	_, err := db.SQLDB.Exec("CREATE TABLE IF NOT EXISTS btc_price(price FLOAT,timestamp INT)")
	if err != nil {
		fmt.Println(err.Error())
	}
}

// InsertNewPriceQuote into database
func (db *DB) InsertNewPriceQuote(price float64, timestamp int64) error {
	stmtIns, err := db.SQLDB.Prepare("INSERT INTO btc_price VALUES(?,?)")
	_, err = stmtIns.Exec(price, timestamp)

	return err
}

// GetLastPrice get the last price from DB
func (db *DB) GetLastPrice() (float64, int64, error) {
	q := "SELECT * from btc_price ORDER BY timestamp DESC LIMIT 1"

	lastPriceStmt, err := db.SQLDB.Prepare(q)
	if err != nil {
		return 0, 0, err
	}

	var lastPrice float64
	var time int64
	err = lastPriceStmt.QueryRow().Scan(&lastPrice, &time)
	if err != nil {
		return 0, 0, err
	}

	return lastPrice, time, nil
}

// GetPriceAtTime returns the price with a given time
// In case of no records at the given time, price with nearest previous timestamp will be returned
// In case of no records prior to a given time, the first record will be returned
func (db *DB) GetPriceAtTime(timestamp string) (float64, int64, error) {
	q := "SELECT * from btc_price WHERE timestamp <= ? ORDER BY timestamp DESC LIMIT 1"

	priceAtTimeStmt, err := db.SQLDB.Prepare(q)
	if err != nil {
		return 0, 0, err
	}

	var price float64
	var time int64
	t, err := helpers.TimeHkToUtc(timestamp)
	if err != nil {
		return 0, 0, err
	}

	err = priceAtTimeStmt.QueryRow(t.Unix()).Scan(&price, &time)
	if err != nil {
		if err == sql.ErrNoRows {
			// In case of no records prior to a given time, the first record will be returned
			return db.getFirstAvailablePrice()
		}
		return 0, 0, err
	}

	return price, time, nil
}

// GetAvgPriceInTimeRange returns the average price within a time range
func (db *DB) GetAvgPriceInTimeRange(startTimeStamp, endTimeStamp string) (float64, error) {
	q := "SELECT price from btc_price WHERE timestamp BETWEEN ? AND ?"

	avgPriceStmt, err := db.SQLDB.Prepare(q)
	if err != nil {
		return 0, err
	}

	startTime, err := helpers.TimeHkToUtc(startTimeStamp)
	endTime, err := helpers.TimeHkToUtc(endTimeStamp)
	if err != nil {
		return 0, err
	}

	rows, err := avgPriceStmt.Query(startTime.Unix(), endTime.Unix())
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	prices := []float64{}

	for rows.Next() {
		var price float64
		err = rows.Scan(&price)
		if err != nil {
			return 0, err
		}
		prices = append(prices, price)
	}

	averagePrice := helpers.CalculateAverage(prices)

	return averagePrice, nil
}

func (db *DB) getFirstAvailablePrice() (float64, int64, error) {
	q := "SELECT * from btc_price ORDER BY timestamp ASC LIMIT 1"

	firstPriceStmt, err := db.SQLDB.Prepare(q)
	if err != nil {
		return 0, 0, err
	}

	var price float64
	var time int64

	err = firstPriceStmt.QueryRow().Scan(&price, &time)
	if err != nil {
		return 0, 0, err
	}

	return price, time, nil
}
