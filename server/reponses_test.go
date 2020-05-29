package server

import (
	"testing"

	"bgl/db"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestLastPrice(t *testing.T) {
	testPrice := 8924.74
	testTime := int64(1589268180)

	dbConn, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer dbConn.Close()

	sqlDB := db.DB{dbConn}
	InitRoutes(&sqlDB)

	rows := sqlmock.NewRows([]string{"price", "timestamp"}).
		AddRow(testPrice, testTime)

	mock.ExpectPrepare(`SELECT\s.*from btc_price`).
		ExpectQuery().
		WillReturnRows(rows)

	resp, err := lastPrice()

	assert := assert.New(t)
	assert.Equal("8924.74", resp.LastPrice)
	assert.Equal("2020-05-12 15:23:00 +0800 HKT", resp.Time)
	assert.Nil(err)
}

func TestPriceAtGivenTime(t *testing.T) {
	testPrice := 8924.74
	testTime := int64(1589268180)

	dbConn, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer dbConn.Close()

	sqlDB := db.DB{dbConn}
	InitRoutes(&sqlDB)

	rows := sqlmock.NewRows([]string{"price", "timestamp"}).
		AddRow(testPrice, testTime)

	mock.ExpectPrepare(`SELECT\s.* from btc_price WHERE timestamp <=`).
		ExpectQuery().
		WillReturnRows(rows)

	resp, err := priceAtGivenTime("2020-05-12T15:23:00")

	assert := assert.New(t)
	assert.Equal("8924.74", resp.Price)
	assert.Equal("2020-05-12 15:23:00 +0800 HKT", resp.Time)
	assert.Nil(err)
}

func TestAveragePrice(t *testing.T) {
	testPrice1 := 10000.00
	testPrice2 := 20000.00

	dbConn, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer dbConn.Close()

	sqlDB := db.DB{dbConn}
	InitRoutes(&sqlDB)
	rows := sqlmock.NewRows([]string{"price"}).
		AddRow(testPrice1).
		AddRow(testPrice2)

	mock.ExpectPrepare(`SELECT price from btc_price WHERE timestamp`).
		ExpectQuery().
		WillReturnRows(rows)

	resp, err := averagePrice("2020-05-12T14:20:00", "2020-05-12T15:20:00")

	assert := assert.New(t)
	assert.Equal("15000.00", resp.AvgPrice)
	assert.Equal("2020-05-12T14:20:00 to 2020-05-12T15:20:00", resp.TimeRange)
	assert.Nil(err)
}
