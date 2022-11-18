package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Median of {56, 85, 92} is: ", median([]float64{56, 85, 92}))         //median= 85
	fmt.Println("Median of {56, 85, 92, 99} is: ", median([]float64{56, 85, 92, 99})) //median= 88.5
}

func median(nums []float64) float64 {
	// Pass by Value: work on a copy, don't change the input slice
	vals := make([]float64, len(nums))
	copy(vals, nums)
	sort.Float64s(vals)

	i := len(vals) / 2
	if len(vals)%2 == 1 { //in case the slice has odd number of items
		return vals[i]
	}

	//in case the slice has even number of items
	return (vals[i-1] + vals[i]) / 2
}
