package main

import (
	"fmt"
)

type Car struct {
	Type   string
	FuelIn float64
}

func (c Car) CalculateDistance() float64 {
	return c.FuelIn * 1.5
}

func main() {
	myCar := Car{
		Type:   "Sedan",
		FuelIn: 20,
	}

	distance := myCar.CalculateDistance()
	fmt.Printf("Estimated distance: %.2f km\n", distance)
}
