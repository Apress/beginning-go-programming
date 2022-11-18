package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("This is a sample code. Learn to code properly")
	term := "code"
	matches, err := (grep(r, term))
	if err != nil {
		log.Fatalln(err.Error())
	}
	if matches != nil {
		fmt.Println("Lines containing \"", term, "\" are :")
		fmt.Println(matches)

	} else {
		fmt.Println("Query term is not present")
	}
}

//grep returns lines in r that contain term
func grep(r io.Reader, term string) ([]string, error) {
	var matches []string
	s := bufio.NewScanner(r)
	for s.Scan() {
		if strings.Contains(s.Text(), term) {
			matches = append(matches, s.Text())
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return matches, nil
}
