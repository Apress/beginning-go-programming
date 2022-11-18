package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)

	states["WA"] = "Washington"
	states["NY"] = "New York"
	states["CA"] = "California"

	keys := make([]string, len(states))
	i := 0
	for key := range states {
		keys[i] = key
		i++
	}

	fmt.Println("Order of Keys Before Sorting: ", keys)

	sort.Strings(keys)
	fmt.Println("Order of Keys After Forting: ", keys)
}

/*OUTPUT
Order of Keys Before Sorting: [WA NY CA]//the display order is not
                                                   guaranteed
Order of Keys After Forting: [CA NY WA]//the display order is guaranteed
*/
