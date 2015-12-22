package main

import (
	"fmt"
	// "html"
	"log"
	"encoding/json"
	"net/http"
)

type data [][]struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func getJSON( w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var d data
	err := decoder.Decode(&d)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Got some!\n%v", d[1][73].X)
}

func main() {
	http.HandleFunc("/js", getJSON)

	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
