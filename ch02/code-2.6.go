package main

import (
	"fmt"
)

func main() {
	var name string
	var age int

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Hello ", name)

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)
	fmt.Println("Your age is ", age)

	var anInt int = 5
	var aFloat float64 = 42
	sum := float64(anInt) + aFloat
	fmt.Printf("Sum: %v, Type: %T \n", sum, sum)

}

/*OUTPUT
Enter your name: Maryam
Hello Maryam
Enter your age: 30
Your age is 30*/
