package main

import (
	"fmt"
)

func myMap(mapObject map[int]int) {
	//make() declares as well as initializes the map to 0
	mapObject = make(map[int]int)
}

func myInt(values []int) {
	//make() declares as well as initializes the slice to 0
	values = make([]int, 5)
}

func main() {
	//mapObject is declared but NOT initialized, i.e., its value is nil
	var mapObject map[int]int
	myMap(mapObject)
	fmt.Println("Is the map equal to nil? ", mapObject == nil)

	//intSlice is declared but NOT initialized, i.e., its value is nil
	var intSlice []int
	myInt(intSlice)
	fmt.Println("Is the slice equal to nil? ", intSlice == nil)
}
