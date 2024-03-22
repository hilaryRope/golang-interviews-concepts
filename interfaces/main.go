package main

import "fmt"

type Vehicle struct {
	manufacturer string
	model        string
	fuel         string
}

func (v Vehicle) startEngine() string {
	message := fmt.Sprintf("This is a car with Manufacturer: %s, Model: %s, Fuel: %s", v.manufacturer, v.model, v.fuel)
	fmt.Println(message)
	return message
}

type Car struct {
	Vehicle
	infotainmentFeatures []string
}

func (c Car) startEngine() string {
	message := fmt.Sprintf("This is a car with Manufacturer: %s, Model: %s, Fuel: %s, InfotainmentFeatures: %v", c.manufacturer, c.model, c.fuel, c.infotainmentFeatures)
	fmt.Println(message)
	return message
}

type Drivable interface {
	startEngine() string
}

func driveTo(d Drivable) {
	d.startEngine()
}

func main() {
	v1 := Vehicle{
		manufacturer: "Fiat",
		model:        "Fiat Punto",
		fuel:         "Petrol",
	}

	v1.startEngine()

	c1 := Car{
		Vehicle:              v1,
		infotainmentFeatures: []string{"Touchscreen with Navigation and Voice Control"},
	}

	c1.startEngine()
	driveTo(c1)
	driveTo(v1)

}
