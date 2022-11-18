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

	/* Print the details of Book1 */
	fmt.Printf("Book1 bTitle : %s\n", Book1.bTitle)
	fmt.Printf("Book1 bAuthorName : %s\n", Book1.bAuthorName)
	fmt.Printf("Book1 bSubject : %s\n", Book1.bSubject)
	fmt.Printf("Book1 book_id : %d\n", Book1.book_id)

	/* Print the details of Book1 */
	fmt.Printf("Book2 bTitle : %s\n", Book2.bTitle)
	fmt.Printf("Book2 bAuthorName : %s\n", Book2.bAuthorName)
	fmt.Printf("Book2 bSubject : %s\n", Book2.bSubject)
	fmt.Printf("Book2 book_id : %d\n", Book2.book_id)
}

/*OUTPUT
Book1 bTitle      : The Go Programming Language
Book1 bAuthorName : Alan A. A. Donovan and Brian W. Kernighan
Book1 bSubject    : A complete guide to Go programming
Book1 book_id     : 6495

Book2 bTitle      : The Complete Book of Arts & Crafts
Book2 bAuthorName : Dawn Purney
Book2 bSubject    : The Complete Book of Arts and Crafts of fun activities for children
Book2 book_id     : 6496
*/
