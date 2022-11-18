package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "Hello from Go!"
	file, err := os.Create("./fromString.txt")
	checkError(err)
	defer file.Close()
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("Wrote a file with %v characters\n", length)

}

/* checkError is a function for error checking */
func checkError(err error) {
	if err != nil {
		panic(err)
	}

}
