package main

import (
	"fmt"
)

func main() {
	poodle := Dog{"Poodle", 10, "woff!"}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight: %v \n", poodle.Breed, poodle.Weight)

	poodle.Speak()
	poodle.Sound = "Arf!"
	poodle.Speak()
	poodle.SpeakThreeTimes()
}

//Dog is a struct
type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

func (d Dog) SpeakThreeTimes() {
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}

/*
OUTPUT
{Poodle 10 woff!}
{Breed:Poodle Weight:10 Sound:woff!}
Breed: Poodle
Weight: 10
Woff!
Arf!
Arf! Arf! Arf!
*/
