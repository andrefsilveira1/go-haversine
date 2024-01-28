package haversine

import (
	"math"
)

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	lat1, lon1, lat2, lon2 = toRadians(lat1), toRadians(lon1), toRadians(lat2), toRadians(lon2)

	latDifference := lat2 - lat1
	lonDifference := lon2 - lon1

	a := math.Pow(math.Sin(latDifference/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(lonDifference/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	Radius := 6371.0
	distance := Radius * c

	return distance
}

func toRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180.0)
}

func Calculate(lat1 float64, lon1 float64, lat2 float64, lon2 float64) float64 {
	return haversine(lat1, lon1, lat2, lon2)
}
