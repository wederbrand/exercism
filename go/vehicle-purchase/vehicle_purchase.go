package purchase

import (
	"sort"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	options := []string{option1, option2}
	sort.Strings(options)
	return options[0] + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age <= 3.0 {
		return originalPrice * 0.8
	}

	if age >= 10.0 {
		return originalPrice * 0.5
	}

	return originalPrice * 0.7
}
