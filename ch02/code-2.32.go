package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)
	fmt.Println(states)
	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"
	fmt.Println(states)

	california := states["CA"]
	fmt.Println(california)

	delete(states, "OR")
	states["NY"] = "New York"
	fmt.Println(states)

	for key, value := range states {
		fmt.Printf("%v: %v \n", key, value)
	}

	keys := make([]string, len(states))
	i := 0
	for key := range states {
		keys[i] = key
		i++
	}
	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)

	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}

/*OUTPUT
map[]
map[CA:California OR:Oregon WA:Washington]
California
map[CA:California NY:New York WA:Washington]
WA: Washington
NY: New York
CA: California
[WA NY CA]
[CA NY WA]
California
New York
Washington
*/
