package main

import "fmt"

// define constants
const minHealthLevel = 0
const maxHealthLevel = 100

// define structs
type HealthStatus struct {
	level int
	city string
}

// define function
func CalculateHealth(city string) HealthStatus {
	// calculate health level
	var healthLevel int
	switch city {
		case "New York":
			healthLevel = 50
		case "Chicago":
			healthLevel = 60
		case "San Francisco":
			healthLevel = 80
		default:
			healthLevel = minHealthLevel
	}
	// check for valid health level
	if healthLevel < minHealthLevel || healthLevel > maxHealthLevel {
		healthLevel = minHealthLevel
	}
	// create return HealthStatus struct
	return HealthStatus {
		level: healthLevel,
		city: city,
	}
}

// main entry point
func main() {
	// define cities
	cities := []string{"New York", "Chicago", "San Francisco"}
	// iterate through all cities and calculate health level
	for _, city := range cities {
		healthStatus := CalculateHealth(city)
		fmt.Printf("Health level in %s is %d\n", healthStatus.city, healthStatus.level)
	}
}