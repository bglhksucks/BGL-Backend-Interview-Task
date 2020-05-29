package db

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetLastPrice(t *testing.T) {
	testPrice := 8924.74
	testTime := int64(1589268180)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"price", "timestamp"}).
		AddRow(testPrice, testTime)

	mock.ExpectPrepare(`SELECT\s.*from btc_price`).
		ExpectQuery().
		WillReturnRows(rows)

	sqlDB := DB{db}
	price, priceTime, err := sqlDB.GetLastPrice()

	assert := assert.New(t)
	assert.Equal(testPrice, price)
	assert.Equal(testTime, priceTime)
	assert.Nil(err)
}

func TestGetPriceAtTime(t *testing.T) {
	testPrice := 8924.74
	testTime := int64(1589268170)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"price", "timestamp"}).
		AddRow(testPrice, testTime)

	mock.ExpectPrepare(`SELECT\s.* from btc_price WHERE timestamp <=`).
		ExpectQuery().
		WillReturnRows(rows)

	sqlDB := DB{db}
	price, priceTime, err := sqlDB.GetPriceAtTime("2020-05-12T15:22:00")

	assert := assert.New(t)
	assert.Equal(testPrice, price)
	assert.Equal(testTime, priceTime)
	assert.Nil(err)

	//Test for cases that not row is founded previous of the timestamp
	rows = sqlmock.NewRows([]string{"price", "timestamp"}) //No rows
	rows2 := sqlmock.NewRows([]string{"price", "timestamp"}).
		AddRow(10000.74, testTime)

	mock.ExpectPrepare(`SELECT\s.* from btc_price WHERE timestamp <=`).
		ExpectQuery().
		WillReturnRows(rows)
	mock.ExpectPrepare(`SELECT\s.* from btc_price ORDER BY timestamp ASC`).
		ExpectQuery().
		WillReturnRows(rows2)

	price, priceTime, err = sqlDB.GetPriceAtTime("2020-05-12T15:22:00")

	assert.Equal(10000.74, price)
	assert.Equal(testTime, priceTime)
	assert.Nil(err)
}

func TestGetAvgPriceInTimeRange(t *testing.T) {
	testPrice1 := 10000.00
	testPrice2 := 20000.00

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"price"}).
		AddRow(testPrice1).
		AddRow(testPrice2)

	mock.ExpectPrepare(`SELECT price from btc_price WHERE timestamp`).
		ExpectQuery().
		WillReturnRows(rows)

	sqlDB := DB{db}
	avgPrice, err := sqlDB.GetAvgPriceInTimeRange("2020-05-12T14:22:00", "2020-05-12T15:22:00")

	assert := assert.New(t)
	assert.Equal(15000.00, avgPrice)
	assert.Nil(err)
}
