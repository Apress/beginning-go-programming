package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "Hello from Go!"
	file, err := os.Create("./fromString.txt")
	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("Wrote a file with %v characters\n", length)
	readFile(file.Name())
	defer file.Close()
}

// readFile is a function for reading content of text file
func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	fmt.Println("Text read from file: ", string(data))
}

// checkError is a function for error checking
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
