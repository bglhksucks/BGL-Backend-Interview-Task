package server

import (
	"bgl/helpers"
)

type lastPriceResp struct {
	LastPrice string `json:"lastPrice"`
	Time      string `json:"time"`
}

type priceResp struct {
	Price string `json:"price"`
	Time  string `json:"time"`
}

type avgPriceResp struct {
	AvgPrice  string `json:"avgPrice"`
	TimeRange string `json:"timeRange"`
}

func lastPrice() (*lastPriceResp, error) {
	price, priceTime, err := dbConn.GetLastPrice()
	if err != nil {
		return nil, err
	}

	lastPriceResp := lastPriceResp{
		helpers.ConvertFloat64ToString(price),
		helpers.TimeUtcToHk(priceTime),
	}

	return &lastPriceResp, nil
}

func priceAtGivenTime(timeStamp string) (*priceResp, error) {
	price, priceTime, err := dbConn.GetPriceAtTime(timeStamp)
	if err != nil {
		return nil, err
	}

	priceResp := priceResp{
		helpers.ConvertFloat64ToString(price),
		helpers.TimeUtcToHk(priceTime),
	}

	return &priceResp, nil
}

func averagePrice(startTime, endTime string) (*avgPriceResp, error) {
	avgPrice, err := dbConn.GetAvgPriceInTimeRange(startTime, endTime)
	if err != nil {
		return nil, err
	}

	avgPriceResp := avgPriceResp{
		helpers.ConvertFloat64ToString(avgPrice),
		startTime + " to " + endTime,
	}

	return &avgPriceResp, nil
}
