package main

import (
	"fmt"
)

func main() {
	total := 1
	for total < 5 {
		total += total
		fmt.Println("Total: ", total)
		if total == 5 {
			goto theEnd
		}
	}

theEnd:
	fmt.Println("End of Program")
}
