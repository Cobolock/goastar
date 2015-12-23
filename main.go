package main

import (
	"fmt"
	//"html"
	"encoding/json"
	"log"
	"net/http"
)

type obstacles [][]Point

type Point struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

func getJSON(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var obsts obstacles
	err := decoder.Decode(&obsts)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Got some!\nObjects to simplify: %d\n", len(obsts))
	for k, obst := range obsts {
		obsts[k] = simplify(obst)
		fmt.Printf("\t%d line: %d => %d\n", k, len(obst), len(obsts[k]))
	}
}

func Append(slice []Point, el Point) []Point {
	n := len(slice)
	if n == cap(slice) {
		newSlice := make([]Point, n, 2*n+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]
	slice[n] = el
	return slice
}

func checkCoLinear(curr, next, last Point) bool {
	return (curr.X-next.X)/(curr.Y-next.Y) == (curr.X-last.X)/(curr.Y-last.Y)
}

func simplify(obst []Point) []Point {
	curr, next, last := 0, 1, 2
	var newObst []Point
	for k, point := range obst {
		last = k
		if k == curr || k == next {
			newObst = Append(newObst, obst[k])
			continue
		}
		if check := checkCoLinear(obst[curr], obst[next], point); check {
			fmt.Printf("\t\t%v, %v, %v\n", obst[curr], obst[next], point)
			next = last
		} else {
			newObst = Append(newObst, obst[last])
			curr = next
			next = last
		}
	}
	return newObst
}

func main() {
	http.HandleFunc("/js", getJSON)

	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
