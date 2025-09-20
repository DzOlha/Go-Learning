package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(lines []string) ([]float64, error) {
	var prices = make([]float64, len(lines))
	for index, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, errors.New("converting price to float failed")
		}
		prices[index] = floatPrice
	}
	return prices, nil
}
