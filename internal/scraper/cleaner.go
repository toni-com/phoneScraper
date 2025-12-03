package scraper

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func ParsePrice(priceStr string) (float64, error) {
	if len(priceStr) == 0 {
		return 0, errors.New("empty price string")
	}

	r := regexp.MustCompile("[^0-9,]")
	priceClean := r.ReplaceAllString(priceStr, "")
	priceClean = strings.Replace(priceClean, ",", ".", -1)
	return strconv.ParseFloat(priceClean, 64)
}
