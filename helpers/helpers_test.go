package helpers

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateAverage(t *testing.T) {
	nums := []float64{1.00, 2.00, 3.00, 4.00, 5.00}
	avgValue := CalculateAverage(nums)

	assert := assert.New(t)
	assert.Equal(3.00, avgValue, "expected: %f, got: %f", 3.00, avgValue)
}

func TestConvertStringToFloat64(t *testing.T) {
	floatNum, err := ConvertStringToFloat64("88.01")
	dataType := reflect.TypeOf(floatNum).Name()

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(dataType, "float64", "expected type: %s, got: %s", "float64", dataType)
	assert.Equal(88.010000, floatNum, "expected type: %f, got: %f", 81.010000, floatNum)

	floatNum, err = ConvertStringToFloat64("ab889")
	assert.NotNil(err)

}

func TestConvertFloat64ToString(t *testing.T) {
	floatNumString := ConvertFloat64ToString(88.01)
	dataType := reflect.TypeOf(floatNumString).Name()

	assert := assert.New(t)
	assert.Equal(dataType, "string", "expected type: %s, got: %s", "string", dataType)
	assert.Equal("88.01", floatNumString, "expected type: %f, got: %f", "88.01", floatNumString)
}

func TestTimeHkToUtc(t *testing.T) {
	utcTime, err := TimeHkToUtc("2020-05-12T15:23:00")

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(int64(1589268180), utcTime.Unix())

	utcTime, err = TimeHkToUtc("2020-0512T15:23:00")
	assert.NotNil(err)
}

func TestTimeUtcToHk(t *testing.T) {
	hkTime := TimeUtcToHk(int64(1589268180))

	assert := assert.New(t)
	assert.Equal("2020-05-12 15:23:00 +0800 HKT", hkTime)
}
