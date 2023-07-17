package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (tUnit *TemperatureUnit) String() string {
	if *tUnit == Celsius {
		return "°C"
	} else {
		return "°F"
	}
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (temp *Temperature) String() string {
	return fmt.Sprintf("%d %s", temp.degree, temp.unit.String())
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (sUnit *SpeedUnit) String() string {
	if *sUnit == MilesPerHour {
		return "mph"
	} else {
		return "km/h"
	}
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (speed *Speed) String() string {
	return fmt.Sprintf("%d %s", speed.magnitude, speed.unit.String())
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (meteorD *MeteorologyData) String() string {
	return fmt.Sprintf("%s: %s, Wind %s at %s, %d%% Humidity", meteorD.location, meteorD.temperature.String(), meteorD.windDirection, meteorD.windSpeed.String(), meteorD.humidity)
}
