package main

import "fmt"

func main() {
	/* Declare and Initialize a Slice of float type */
	marks := []float32{10, 12.6, 20.0, 37.56, 48.74, 50.0, 64.15,
		79.63, 8.75}

	/*pass slice to user-defined function to print it details*/
	printSliceDetails(marks)

	/* Printing the Elements of the Original Slice */
	fmt.Println("Original Slice =", marks)

	/* Printing sub-slice of marks slice beginning from
	   index 1(inclusive) to index 5(excluded)*/
	fmt.Println("Marks[1:5] = ", marks[1:5])

	/*  Not indicating the lower bound inferred as 0 */
	fmt.Println("Marks[:4] =", marks[:4])

	/* Not indicating the upper bound is inferred as len(slice) */
	fmt.Println("Marks[3:] =", marks[3:])
	marks1 := make([]float32, 0, 5)
	/*pass slice to user-defined function to print it details*/
	printSliceDetails(marks1)

	/* Printing sub-slice of marks beginning from
	   index 0(inclusive) to index 3(exclusive) */
	marks2 := marks[:3] //storing subslice in slice named marks2
	/*pass slice to user-defined function to print it details*/
	printSliceDetails(marks2)

	/* Printing sub-slice of marks beginning from
	   index 4(inclusive) to index 5(exclusive) */
	marks3 := marks[4:5]
	/*pass slice to user-defined function to print it details*/
	printSliceDetails(marks3)
}

func printSliceDetails(x []float32) {
	fmt.Printf("Length=%d Capacity=%d Slice=%v\n", len(x), cap(x), x)
}

/*OUTPUT
Length=9 Capacity=9 Slice=[10 12.6 20 37.56 48.74 50 64.15 79.63 8.75]
Original Slice = [10 12.6 20 37.56 48.74 50 64.15 79.63 8.75]
Marks[1:5] =  [12.6 20 37.56 48.74]
Marks[:4] = [10 12.6 20 37.56]
Marks[3:] = [37.56 48.74 50 64.15 79.63 8.75]
Length=0 Capacity=5 Slice=[]
Length=3 Capacity=9 Slice=[10 12.6 20]
Length=1 Capacity=5 Slice=[48.74]
*/
