package main

import "fmt"

func main() {

	/* Declaring a slice of int type */
	var nums []int

	/* pass slice to user-defined function to print details */
	printSliceDetails(nums)

	/* Nil slice are allowed with Append */
	nums = append(nums, 10)
	printSliceDetails(nums)

	/* Adding one element to the slice */
	nums = append(nums, 100)
	printSliceDetails(nums)

	/* At one time adding more than one element to the slice */
	nums = append(nums, 1000, 10000, 100000)
	printSliceDetails(nums)

	/* Creating a slice nums1 that has double the capacity of nums slice*/
	nums1 := make([]int, len(nums), (cap(nums))*2)

	/* Copy the elements stored in nums into nums1 */
	copy(nums1, nums)

	/* pass slice to user-defined function to print details */
	printSliceDetails(nums1)

	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println("Before: ", colors)
	colors = append(colors[1:len(colors)])
	fmt.Println("Items after removig 1st element:", colors)

}

func printSliceDetails(x []int) {
	fmt.Printf("Length=%d Capacity=%d Slice=%v\n", len(x), cap(x), x)
}

/*OUTPUT
Length=0 Capacity=0 Slice=[]
Length=1 Capacity=1 Slice=[10]
Length=2 Capacity=2 Slice=[10 100]
Length=5 Capacity=6 Slice=[10 100 1000 10000 100000]
Length=5 Capacity=12 Slice=[10 100 1000 10000 100000]
*/
