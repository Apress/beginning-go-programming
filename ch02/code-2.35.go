package main

import "fmt"

type Books struct {
	bTitle      string
	bAuthorName string
	bSubject    string
	book_id     int
}

func main() {
	var Book1 Books /* Declaring variable of type Book */
	var Book2 Books /* Declaring variable of type Book */

	/* Access member fields of Book struct to define Book1 */
	Book1.bTitle = "The Go Programming Language"
	Book1.bAuthorName = "Alan A. A. Donovan and Brian W. Kernighan"
	Book1.bSubject = "A complete guide to Go programming"
	Book1.book_id = 6495

	/* Access member fields of Book struct to define Book2 */
	Book2.bTitle = "The Complete Book of Arts & Crafts"
	Book2.bAuthorName = "Dawn Purney"
	Book2.bSubject = "The Complete Book of Arts and Crafts of fun activities for children"
	Book2.book_id = 6496

	/* Print the details of Book1 by passing it as an argument to function */
	printBookDetails(Book1)

	/* Print the details of Book2 by passing it as an argument to function */
	printBookDetails(Book2)
}

func printBookDetails(book Books) {
	fmt.Printf("\nTitle : %s\n", book.bTitle)
	fmt.Printf("Authors : %s\n", book.bAuthorName)
	fmt.Printf("Subject : %s\n", book.bSubject)
	fmt.Printf("Book ID : %d\n", book.book_id)
}

/*OUTPUT
Title   : The Go Programming Language
Authors  : Alan A. A. Donovan and Brian W. Kernighan
Subject : A complete guide to Go programming
Book ID : 6495

Title    : The Complete Book of Arts & Crafts
Authors  : Dawn Purney
Subject : The Complete Book of Arts and Crafts of fun activities for children
Book ID : 6496
*/
