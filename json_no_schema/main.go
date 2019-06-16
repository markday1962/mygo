package main

import (
	"encoding/json"
	"fmt"
)

var JSON = `{
"firstName": "Mark",
"lastName": "Day",
"age": 57,
"address": [
	{
		"HouseNumber": "128a",
		"Street": "Columbia Road",
		"City": "London",
		"Postcode": "E2 7RG"
	}
]
}`

func main() {
	var f interface{}
	err := json.Unmarshal([]byte(JSON), &f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)
	//map[address:[map[City:London HouseNumber:128a Postcode:E2 7RG Street:Columbia Road]] age:57 firstName:Mark lastName:Day]
	printJSON(f)
}

func printJSON(v interface{}) {
	fmt.Println(v)
}
