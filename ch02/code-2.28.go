package main

import (
	"fmt"
	"sort"
)

func main() {
	intSlice := make([]int, 5)
	intSlice[0] = 6
	intSlice[1] = 99
	intSlice[2] = 45
	intSlice[3] = 34
	intSlice[4] = 1

	fmt.Println("Original Slice: ", intSlice)

	sort.Ints(intSlice)

	fmt.Println("Sorted Slice: ", intSlice)
}

/*OUTPUT
Original Slice:  [6 99 45 34 1]
Sorted Slice: [1 6 34 45 99]
*/
