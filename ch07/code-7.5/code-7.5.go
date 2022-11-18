package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/mux"
)

var (
	db     = make(map[string][]byte)
	dbLock sync.RWMutex
)

const maxSize = 1 << 20 //MB
//client side
const apiBase = "http://localhost:8080/kv"

func list() error {
	resp, err := http.Get(apiBase)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %d %s", resp.StatusCode, resp.Status)
	}

	defer resp.Body.Close()
	var keys []string
	if json.NewDecoder(resp.Body).Decode(&keys); err != nil {
		return err
	}

	for _, key := range keys {
		fmt.Println(key)
	}
	return nil
}

func set(key string) error {
	url := fmt.Sprintf("%s/%s", apiBase, key)
	resp, err := http.Post(url, "application/octet-stream", os.Stdin)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Bad Status: %d %s", resp.StatusCode, resp.Status)
	}

	var reply struct {
		Key  string
		Size int
	}

	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&reply); err != nil {
		return err
	}
	fmt.Printf("%s: %d bytes\n", reply.Key, reply.Size)
	return nil
}

func get(key string) error {
	url := fmt.Sprintf("%s/%s", apiBase, key)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Bad Status: %d %s", resp.StatusCode, resp.Status)
	}

	_, err = io.Copy(os.Stdout, resp.Body)
	return err
}

//server side
func handleSet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	defer r.Body.Close()
	rdr := io.LimitReader(r.Body, maxSize)
	data, err := ioutil.ReadAll(rdr)
	if err != nil {
		log.Printf("read error: %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dbLock.Lock()
	defer dbLock.Unlock()
	db[key] = data

	resp := map[string]interface{}{
		"key":  key,
		"size": len(data),
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("error sending: %s", err)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	dbLock.RLock()
	defer dbLock.RUnlock()

	data, ok := db[key]
	if !ok {
		log.Printf("error get - unknown key: %q", key)
		http.Error(w, fmt.Sprintf("%q not found", key), http.StatusNotFound)
		return
	}

	if _, err := w.Write(data); err != nil {
		log.Printf("error sending: %s", err)
	}
}

func handleList(w http.ResponseWriter, r *http.Request) {
	dbLock.RLock()
	defer dbLock.RUnlock()

	keys := make([]string, 0, len(db))
	for key := range db {
		keys = append(keys, key)
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(keys); err != nil {
		log.Printf("error sending: %s", err)
	}
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: kv get|set|list [key]")
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatalf("error: wrong  number of arguments")
	}

	switch flag.Arg(0) {
	case "get":
		key := flag.Arg(1)
		if key == "" {
			log.Fatal("error: missing key")
		}
		if err := get(key); err != nil {
			log.Fatal(err)
		}
	case "set":
		key := flag.Arg(1)
		if key == "" {
			log.Fatal("error: missing key")
		}
		if err := set(key); err != nil {
			log.Fatal(err)
		}
	case "list":
		key := flag.Arg(1)
		if key == "" {
			log.Fatal("error: missing key")
		}
		if err := list(); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("error: unknown command: %q", flag.Arg(0))
	}

	r := mux.NewRouter()
	r.HandleFunc("/kv/{key}", handleSet).Methods("POST")
	r.HandleFunc("/kv/{key}", handleGet).Methods("GET")
	r.HandleFunc("/kv", handleList).Methods("GET")
	http.Handle("/", r)

	addr := ":8080"
	log.Printf("Server Ready On %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
