package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

/* Example data in file to read
1542784314:0;git push
1542784378:0;ls
1542784308:0;go test
1542784310:0;go run -v
1542784311:0;go test -v
*/

var cmdRe = regexp.MustCompile(`;go ([a-z]+)`)

//cmdFreq returns the frequency of the go subcommand from the given file
func cmdFreq(fileName string) (map[string]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	freqs := make(map[string]int)
	s := bufio.NewScanner(file)
	for s.Scan() {
		matches := cmdRe.FindStringSubmatch(s.Text())
		if len(matches) == 0 {
			continue
		}
		cmd := matches[1]
		freqs[cmd]++
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return freqs, nil
}

func main() {
	result, _ := cmdFreq("./sample.txt")
	for key, _ := range result {
		fmt.Println("Subcommand: ", key, " Count: ", result[key])
	}
}

/*
OUTPUT
Subcommand:  test  Count:  2
Subcommand:  run  Count:  1
*/
