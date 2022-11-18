package main

import (
	"fmt"
)

func main() {
	//Explicit Declaration
	var aStringVariable string = "i am a string"
	fmt.Println(aStringVariable)
	fmt.Printf("Printing variable along text: %s \n", aStringVariable)

	var anotherStringVariable string
	fmt.Println(anotherStringVariable)
	fmt.Printf("Printing variable along text: %s \n", anotherStringVariable)

	var defaultInt int
	fmt.Println(defaultInt)

	//Implicit Declaration
	myString := "Implicit Declaration of string"
	fmt.Println(myString)
}
