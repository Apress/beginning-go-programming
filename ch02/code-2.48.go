package main

import (
	"fmt"
)

func main() {
	total := addValues(40, 50, 60)
	fmt.Println("Sum of Passed Values: ", total)
}

func addValues(values ...int) int {
	sum := 0
	for _, val := range values {
		sum += val
	}
	return sum
}

/*
OUTPUT
Sum of Passed Values:  150
*/
