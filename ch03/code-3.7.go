package main

import (
	"fmt"
)

func main() {
	values := []int{22, 85, 36, 94, 50, 67, 17, 18}
	fmt.Println("Before Filtering: ", values)
	fmt.Println("After Filtering: ", filter(isOdd, values))
}

//filter returns a slice with only the values that pred(val) returned true
func filter(pred func(int) bool, values []int) []int {
	var out []int
	for _, val := range values {
		if pred(val) {
			out = append(out, val)
		}
	}
	return out
}

func isOdd(n int) bool {
	return n%2 == 1
}

/*
OUTPUT
[85 67 17]
*/
