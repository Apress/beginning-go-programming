package main

import "fmt"

func main() {
	var marks []float64      //declaring slice of float type
	printItemsOfSlice(marks) //pass slice to a function

	if marks == nil {
		fmt.Printf("Slice is Nil")
	}
}

// function that accepts a slice and prints its details
func printItemsOfSlice(x []float64) {
	fmt.Printf("Length=%d Capacity=%d Slice=%v\n", len(x), cap(x), x)
}

/*OUTPUT
Length=0 Capacity=0 Slice=[]
Slice is Nil
*/
