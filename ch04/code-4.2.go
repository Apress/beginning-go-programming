package main

import (
	"fmt"
	"strings"
)

type Letter struct {
	Symbol, English string
}

var letters = []Letter{
	{"Σ", "Sigma"},
}

func englishFor(greek string) (string, error) {
	for _, letter := range letters {
		if strings.EqualFold(greek, letter.Symbol) {
			return letter.English, nil
		}
	}
	return "", fmt.Errorf("Unknown Greek Letter: %#v", greek)
}
func main() {
	fmt.Println(englishFor("Σ"))
	fmt.Println(englishFor("σ"))
	fmt.Println(englishFor("ς"))
	fmt.Println(englishFor("�"))
}

/*
OUTPUT
Sigma <nil>
Sigma <nil>
Sigma <nil>
 Unknown Greek Letter: "�"
*/
