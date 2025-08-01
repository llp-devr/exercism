// Package weather provides functionalities related to weather forecasting.
package weather


// CurrentCondition holds the current weather condition.
var CurrentCondition string

// CurrentLocation holds the current location for which the weather is relevant.
var CurrentLocation string


// Forecast sets the current location and condition based on the provided city and condition.
// It returns a string that represents a weather forecast message.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
