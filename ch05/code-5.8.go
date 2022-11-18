package main

import (
	"fmt"
)

//A thermostat measures and controls the temperature
type Thermostat struct {
	id    string
	value float64
}

//Value return the current temperature in Celsius
func (t *Thermostat) Value() float64 {
	return t.value
}

//ID returns the thermostat ID
func (t *Thermostat) ID() string {
	return t.id
}

//Set tells the thermostat to set the temperature
func (t *Thermostat) Set(value float64) {
	t.value = value
}

//Kind returns the device kind
func (t *Thermostat) Kind() string {
	return "thermostat"
}

//Camera is a security camera
type Camera struct {
	id string
}

//ID return the camera ID
func (c *Camera) ID() string {
	return c.id
}

func (*Camera) Kind() string {
	return "camera"
}

type Sensor interface {
	ID() string
	Kind() string
}

func printAll(sensors []Sensor) {
	for _, s := range sensors {
		fmt.Printf("%s <%s>\n", s.ID(), s.Kind())
	}
}

func main() {
	t := Thermostat{"Living Room", 16.2}
	c := Camera{"Baby Room"}

	sensors := []Sensor{&t, &c}
	printAll(sensors)
	/*fmt.Printf("%s Before: %.2f\n", t.ID, t.Value())
	t.Set(18)
	fmt.Printf("%s After: %.2f\n", t.ID, t.Value())*/
}

/*
OUTPUT
Living Room <thermostat>
Baby Room <camera>
*/
