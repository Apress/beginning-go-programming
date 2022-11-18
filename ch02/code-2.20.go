package main

import (
	"fmt"
)

func main() {
	/*declaring and initializing an array "array1" of size 3 and type as
	  floating-point*/
	var array1 = []float32{10.5, 5.2, 2.88}

	var array2 [10]int
	var i, j int

	//Initializing elements of the array
	for i = 0; i < 10; i++ {
		array2[i] = i + 50 //setting element at location i to i+50
	}

	//Print the value of each elements of array1
	fmt.Println("Elements stored in Array1")
	for j = 0; j < 3; j++ {
		fmt.Printf("Element[%d] = %f \n", j, array1[j])
	}

	//Print the value of each elements of array2
	fmt.Println("Elements stored in Array1")
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d \n", j, array2[j])
	}

	fmt.Println("Size of array1: ", len(array1))
	fmt.Println("Size of array2: ", len(array2))

}

/*OUTPUT
$ go run array-1.go
Elements stored in Array1
Element[0] = 10.500000
Element[1] = 5.200000
Element[2] = 2.880000
Elements stored in Array1
Element[0] = 50
Element[1] = 100
Element[2] = 150
Element[3] = 200
Element[4] = 250
Element[5] = 300
Element[6] = 350
Element[7] = 400
Element[8] = 450
Element[9] = 500 */
