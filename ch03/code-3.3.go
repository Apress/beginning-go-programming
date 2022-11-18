package main

import (
	"fmt"
)

var moby = []string{
	"shall", "pass", "shall", "away", "too"}

func main() {
	fmt.Println(frequency(moby))
}

func frequency(words []string) map[string]int {
	freq := make(map[string]int)
	for _, w := range words {
		freq[w]++
	}
	return freq
}
