// Package weather provides tools relating to the weather.
package weather

// CurrentCondition represents the conditions the weather currently is.
var CurrentCondition string

// CurrentLocation represents the location the weather is currently in.
var CurrentLocation string

// Forecast returns the forecast for a given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
