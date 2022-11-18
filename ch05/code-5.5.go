package main

import (
	"fmt"
)

//A thermostat measures and controls the temperature
type Thermostat struct {
	ID    string
	value float64
}

//Value return the current temperature in Celsius
func (t *Thermostat) Value() float64 {
	return t.value
}

//Set tells the thermostat to set the temperature
func (t *Thermostat) Set(value float64) {
	t.value = value
}

//Kind returns the device kind
func (t *Thermostat) Kind() string {
	return "thermostat"
}

func main() {
	t := Thermostat{"Living Room", 16.2}
	fmt.Printf("%s Before: %.2f\n", t.ID, t.Value())
	t.Set(18)
	fmt.Printf("%s After: %.2f\n", t.ID, t.Value())
}

/*
OUTPUT
Living Room Before: 16.20
Living Room After: 18.00
*/
