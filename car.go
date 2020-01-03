package main

import (
	"io"
)

//Car Structure
type Car struct {
	Company  string
	Model    string
	Color    string
	Tires    int
	Mileage  int
	TopSpeed int
	Years    int
	Output   io.Writer
}

//NewCar created
func NewCar(output io.Writer) *Car {
	return &Car{
		Company:  "Nissan",
		Model:    "Skyline GTR R34",
		Color:    "Black",
		Tires:    0,
		Mileage:  0,
		TopSpeed: 160,
		Years:    0,
		Output:   output,
	}
}

//satisfies the car interface

//Age allows the car to get older
func (c *Car) Age() {
	c.Years++
}

//Miles adds mileage to car
func (c *Car) Miles() {
	c.Mileage++
}

//Drive makes the car drive
func (c *Car) Drive() {
	c.Output.Write([]byte("Bruce goes for a drive....pedal to the metal<br>\n"))
	c.Mileage = c.Mileage + 1000
}

//Oil changes the oil in the car
func (c *Car) Oil() {
	c.Output.Write([]byte("Bruce gets the oil changed...keep it fresh broski<br>\n"))
	c.Mileage = 0

}

//Park keeps the car parked
func (c *Car) Park() {
	c.Output.Write([]byte("It is raining, ain't no way Bruce is driving this gem in that<br>\n"))
	c.Mileage = c.Mileage + 1000
}

//Wash washes the Skyline
func (c *Car) Wash() {
	c.Output.Write([]byte("Car is dirty, time for a scrub!<br>\n"))
	c.Mileage = c.Mileage + 1000

}

//Drift the car
func (c *Car) Drift() {
	c.Output.Write([]byte("Drift time....getting slidewayzzz<br>\n"))
	c.Tires = c.Tires + 1
	c.Mileage = c.Mileage + 1000

}
