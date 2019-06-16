package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Age   int    `json:"age"`
}

//example of a json representation that is of type string
var JSON = `{
	"first": "Mark",
	"last": "Day",
	"age": 57
}`

func main() {
	var p Person
	err := json.Unmarshal([]byte(JSON), &p) //parse json data into the p instance of type Person
	if err != nil {
		fmt.Println("Json unmarshal error", err)
	}
	fmt.Println(p)
	fmt.Println("Forename:", p.First, ",Surname:", p.Last, ",Age:", p.Age)
}
