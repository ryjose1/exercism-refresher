// Package weather provides tools to aid weather forcasting programs
package weather

// CurrentCondition is the weather description at the present moment for a particular forecast
var CurrentCondition string

// CurrentLocation is the place that a particular forecast is referencing
var CurrentLocation string

// Forecast returns a string that is a human-readable description of the current weather condition at a location
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
