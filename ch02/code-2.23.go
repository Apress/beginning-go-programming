package main

import (
	"fmt"
)

func main() {
	var marks = make([]float64, 3, 5) //declaring slice of length 3 and capacity 5
	printItemsOfSlice(marks)          //pass slice to a function
}

// function that accepts a slice and prints its details
func printItemsOfSlice(x []float64) {
	fmt.Printf("Length=%d Capacity=%d Slice=%v\n", len(x), cap(x), x)
}

/*output
s
*/
