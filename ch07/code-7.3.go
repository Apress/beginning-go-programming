package main

import (
	"encoding/json"
	"gorilla/mux"
	"io"
	"log"
	"net/http"
	"time"
)

// DB is a database
type DataBase struct{}

// Add adds a metric to the database
func (db *DataBase) Add(m Metric) string {
	return "success"
}

type Metric struct {
	Time   time.Time `json:"time"`
	Host   string    `json:"host"`
	CPU    float64   `json:"cpu"`    //CPU load
	Memory float64   `json:"memory"` //MB
}

func handleMetric(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "This Method is Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var db *DataBase
	defer r.Body.Close()
	var m Metric
	const maxSize = 1 << 20 //MB
	dec := json.NewDecoder(io.LimitReader(r.Body, maxSize))
	if err := dec.Decode(&m); err != nil {
		log.Printf("Error Decoding: %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := db.Add(m)
	log.Printf("metric: %+v (id=%s)", m, id)

	w.Header().Set("Content-Type", "application/json")
	resp := map[string]interface{}{
		"id": id,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("error reply: %s", err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/metrics.json", handleMetric).Methods("POST")

	http.Handle("/", r)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
