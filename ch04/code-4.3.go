package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	/* Checking directly if the regex pattern matches a string */
	match, _ := regexp.MatchString("p([a-z]+)ch", "preach")
	fmt.Println(match)

	/* Compiling an Optimized Regexp struct to get a Regexp object
	   that can be used to match against text */
	r, _ := regexp.Compile("p([a-z]+)ch")

	/* MatchString checks if passed string has any match of the
	   regular expression */
	fmt.Println(r.MatchString("preach"))

	/* FindString checks if the passed string holds any text
	   matching leftmost text of the regular expression */
	fmt.Println(r.FindString("preach patch"))

	/* FindStringIndex finds the first match in the passed string
	   that matches the regular expression and returns the
	   start and end indexes of match instead text */
	fmt.Println("Start and End Indexes of Match:",
		r.FindStringIndex("pinch pouch"))

	/* FindStringSubmatch finds the leftmost match in the passed
	   string that matches the regular expression and the submatch */
	fmt.Println(r.FindStringSubmatch("poach pitch"))

	/* FindStringSubmatch finds the leftmost match in the passed string
	   that matches the regular expression and the submatch and returns
	   the start and end index instead of text. Here, the end index for
	   match is exclusive. */
	fmt.Println(r.FindStringSubmatchIndex("punch"))

	/*The All variants finds all matches in passed string*/
	/* Find all matches in a given input for a regular expression*/
	fmt.Println(r.FindAllString("parch patch pitch", -1))

	/* FindAllStringSubmatchIndex returns a slice containing all of the
	   matches regular of the expression*/
	fmt.Println("Indexes of All Matches and Submatches:", r.FindAllStringSubmatchIndex(
		"potch pooch porch", -1))

	/* Limiting the number of returned matches by passing a non-negative
	   integer as an argument */
	fmt.Println(r.FindAllString("prelaunch postlaunch pitch", 2))

	/* Checking if the byte slice contains text matching to
	   regular expression*/
	fmt.Println(r.Match([]byte("pinch")))

	/* Using MustCompile with global variables*/
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	/* Replacing string subsets with other values */
	fmt.Println(r.ReplaceAllString("pinch it!", "hurt"))

	/* Transforming matched text with a specified function.*/
	in := []byte("The prelaunch")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}

/*
OUTPUT
true
true
preach
Start and End Indexes of Match: [0 5]
[poach oa]
[0 5 1 3]
[parch patch pitch]
Indexes of All Matches and Submatches: [[0 5 1 3] [6 11 7 9]
[12 17 13 15]]
[prelaunch postlaunch]
true
regexp: p([a-z]+)ch
hurt it!
The PRELAUNCH
*/
