// Package cars provides tools for analyzing the productivity of a car assembly line
package cars

import (
	"log"
	"math"
)

// DefaultCarsPerHour is the rate of cars the assembly line is expected to produce at speed 1
const DefaultCarsPerHour = 221

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
	switch speed {
	case 0:
		return 0.0
	case 1, 2, 3, 4:
		return 1.0
	case 5, 6, 7, 8:
		return 0.9
	case 9, 10:
		return .77
	default:
		log.Print("Success rate not known for given speed, will assume no defects")
		return 1.0
	}
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	return float64(speed) * float64(DefaultCarsPerHour) * SuccessRate(speed)
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	return int(CalculateProductionRatePerHour(speed) / 60.0)
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	return math.Min(CalculateProductionRatePerHour(speed), limit)
}
