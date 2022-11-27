package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "truck" || kind == "car" {
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	bestOption := option1
	if option1 > option2 {
		bestOption = option2
	}
	return fmt.Sprintf("%s is clearly the better choice.", bestOption)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var val float64 = 0.80
	if age >= 10 {
		val = 0.50
	} else if age >= 3 {
		val = 0.70
	}
	return originalPrice * val
}
