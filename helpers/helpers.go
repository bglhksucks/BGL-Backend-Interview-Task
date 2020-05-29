package helpers

import (
	"strconv"
	"time"
)

// CalculateAverage calculates the average value in a price array
func CalculateAverage(prices []float64) float64 {
	total := float64(0)
	for _, number := range prices {
		total = total + number
	}
	average := total / float64(len(prices))

	return average
}

// ConvertStringToFloat64 converts string to float64
func ConvertStringToFloat64(num string) (float64, error) {
	floatInString, err := strconv.ParseFloat(num, 64)

	return floatInString, err
}

// ConvertFloat64ToString converts float64 to string
func ConvertFloat64ToString(num float64) string {
	floatInString := strconv.FormatFloat(num, 'f', 2, 64)

	return floatInString
}

// TimeHkToUtc converts a Hong Kong Timestamp into UTC Time
func TimeHkToUtc(timeStamp string) (time.Time, error) {
	// HK is UTC +8
	t, err := time.Parse(time.RFC3339, timeStamp+"+00:00") //eg. 2020-05-12T15:23:00
	utcTime := t.Add(time.Hour * -8)

	return utcTime, err
}

// TimeUtcToHk converts a Unix Utc Timestamp into HK Time
func TimeUtcToHk(utcTimeUnix int64) string {
	utcTime := time.Unix(utcTimeUnix, 0)

	loc, _ := time.LoadLocation("Asia/Hong_Kong")
	hkTime := utcTime.In(loc)

	// hkTime := utcTime.Add(time.Hour * 8)

	return hkTime.String()
}
