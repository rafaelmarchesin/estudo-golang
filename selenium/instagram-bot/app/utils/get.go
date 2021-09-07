package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func GetRandomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func GetRandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func GetRandomTimeDuration(min, max int) time.Duration {
	r := rand.Intn(max-min) + min
	return time.Duration(r)
}

func GetFormatedNum(s string) (int, error) {
	s = strings.Replace(s, ".", "", -1)
	s = strings.Replace(s, ",", "", -1)
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}
