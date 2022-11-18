package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

/*{
	"lastCheckTime": "2022-02-15 09:20:15 PM",
	"stations":[
		{
			"id":3,
			"name": "station 3",
			"status": "Active",
			"statusKey": 1,
			"lastCheck": {
				"time": "2016-01-22 04:30:15 PM",
				"success": true
			}
		},
		{
			"id":1,
			"name": "station 2",
			"status": "Active",
			"lastCheck": {
				"time": "2020-01-22 04:32:41 PM",
				"success": true
			}
		},

	]

}
*/
//laggingStations return stations that are lagging in their check time
func laggingStations(r io.Reader, timeout time.Duration) ([]string, error) {
	var reply struct {
		LastCheckTime string
		Stations      []struct {
			Name      string
			Status    string
			LastCheck struct {
				Time string
			}
		}
	}

	dec := json.NewDecoder(r)
	if err := dec.Decode(&reply); err != nil {
		return nil, err
	}

	checkTime, err := parseTime(reply.LastCheckTime)
	if err != nil {
		return nil, err
	}

	var lagging []string
	for _, station := range reply.Stations {
		if station.Status != "Active" {
			continue
		}
		lastCheck, err := parseTime(station.LastCheck.Time)
		if err != nil {
			return nil, err
		}
		if checkTime.Sub(lastCheck) > timeout {
			lagging = append(lagging, station.Name)
		}
	}
	return lagging, nil
}

func parseTime(ts string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05 PM", ts)
}

func main() {
	file, err := os.Open("stations.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lagging, err := laggingStations(file, time.Minute)
	if err != nil {
		log.Fatal(err)
	}

	for _, name := range lagging {
		fmt.Println(name)
	}
}
/*
OUTPUT
station 3
*/