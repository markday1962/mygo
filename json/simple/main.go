package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

//example of a json representation that is of type string
var JSON = `{
	"firstName": "Mark",
	"lastName": "Day",
	"age": 57
}`

func main() {
	var p Person
	err := json.Unmarshal([]byte(JSON), &p) //parse json data into the p instance of type Person
	if err != nil {
		fmt.Println("Json unmarshal error", err)
	}
	fmt.Println("Printing p:", p)
	fmt.Println("Forename:", p.FirstName, ",Surname:", p.LastName, ",Age:", p.Age)
}
