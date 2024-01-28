package main

import (
	"fmt"
	"math"
)

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	// Convert latitude and longitude from degrees to radians
	lat1, lon1, lat2, lon2 = toRadians(lat1), toRadians(lon1), toRadians(lat2), toRadians(lon2)

	// Compute differences
	dlat := lat2 - lat1
	dlon := lon2 - lon1

	// Haversine formula
	a := math.Pow(math.Sin(dlat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dlon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// Radius of the Earth in kilometers
	R := 6371.0

	// Calculate distance
	distance := R * c

	return distance
}

func toRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180.0)
}

func main() {
	// Example usage
	lat1, lon1 := 37.7749, -122.4194 // San Francisco, CA
	lat2, lon2 := 34.0522, -118.2437 // Los Angeles, CA

	distance := haversine(lat1, lon1, lat2, lon2)
	fmt.Printf("The distance between the two points is %.2f km.\n", distance)
}
