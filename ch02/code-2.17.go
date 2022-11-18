package main

import (
	"fmt"
)

func main() {
	value1 := 42
	var pointer1 = &value1
	fmt.Println("Value of pointer1: ", *pointer1)

	value2 := 42.13
	pointer2 := &value2
	fmt.Println("Value1: ", *pointer2)

	value3 := 32.5
	pointer3 := &value3
	*pointer3 = *pointer3 / 31
	fmt.Println("Pointer3: ", *pointer3)
	fmt.Println("Value3: ", value3)

}
