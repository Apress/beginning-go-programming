package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Response Type: %T\n", resp)

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	content := string(bytes)

	tours := toursFromJson(content)

	for _, tour := range tours {
		fmt.Println(tour.Name, "  ", tour.Price)
	}
}

func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20) //slice of Tour array with initial size 0 and capacity 20

	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	if err != nil {
		panic(err)
	}

	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		if err != nil {
			panic(err)
		}
		tours = append(tours, tour)
	}
	return tours
}

type Tour struct {
	Name, Price string
}

/*OUTPUT
2 Days Adrift the Salton Sea   350
A Week of Wine   850
Amgen Tour of California Special   6000
Avila Beach Hot springs   1000
Big Sur Retreat   750
Channel Islands Excursion   150
Coastal Experience   1500
Cycle California: My Way   1200
Day Spa Package   550
Endangered Species Expedition   600
Fossil Tour   500
Hot Salsa Tour   400
Huntington Library and Pasadena Retreat Tour   225
In the Steps of John Muir   600
Joshua Tree: Best of the West Tour   150
Kids L.A. Tour   200
Mammoth Mountain Adventure   800
Matilija Hot springs   1000
Mojave to Malibu   200
Monterey to Santa Barbara Tour   2500
Mountain High Lift-off   800
Olive Garden Tour   75
Oranges & Apples Tour   350
Restoration Package   900
The Death Valley Survivor's Trek   250
The Mt. Whitney Climbers Tour   650
*/
