package haversine

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHaversine(t *testing.T) {
	coordinates := []struct {
		lat1, lon1, lat2, lon2 float64
		want                   float64
	}{
		{37.7749, -122.4194, 34.0522, -118.2437, 559.1205770615534}, // San Francisco, CA to Los Angeles, CA
		{40.7128, -74.0060, 41.8781, -87.6298, 1144.2912739463477},  // New York, NY to Chicago, IL
		{51.5074, -0.1278, 48.8566, 2.3522, 343.5560603410416},      // London, UK to Paris, France
	}

	for _, coord := range coordinates {
		name := fmt.Sprintf("Distance between (%f, %f) and (%f, %f)", coord.lat1, coord.lon1, coord.lat2, coord.lon2)

		t.Run(name, func(t *testing.T) {
			distance := Calculate(coord.lat1, coord.lon1, coord.lat2, coord.lon2)
			assert.Equal(t, coord.want, distance)
		})
	}
}
