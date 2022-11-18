package main

import (
	"fmt"
)

func main() {
	theAnswer := 42
	var result string

	if theAnswer < 0 {
		result = "Less than Zero"
	} else if theAnswer == 0 {
		result = "Equal to Zero"
	} else {
		result = "Greater than Zero"
	}

	fmt.Println(result)
}

/*OUTPUT
Greater than Zero
*/
