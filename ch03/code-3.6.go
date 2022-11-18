package main

import "os"

func main() {

	//Trigger panic with Custom Error Message
	panic("Problem!")

	_, err := os.Create("/tmp/file")

	if err != nil {
		panic(err)
	}
}
