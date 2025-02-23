package utils

import geo "github.com/kellydunn/golang-geo"

// HaversineDistance calculates the great-circle distance between two points
// CalculateDistance returns the great-circle distance between two coordinates
func CalculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	p1 := geo.NewPoint(lat1, lon1)
	p2 := geo.NewPoint(lat2, lon2)
	return p1.GreatCircleDistance(p2)
}
