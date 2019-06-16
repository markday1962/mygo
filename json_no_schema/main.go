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
	simplePrintJSON(f)
	walkJSON(f)
}

func simplePrintJSON(v interface{}) {
	fmt.Println(v)
	m := v.(map[string]interface{})
	// accessing top level data if we know the structure of v
	fmt.Println(m["firstName"])
	fmt.Println(m["lastName"])
	fmt.Println(m["address"])
}

func walkJSON(v interface{}) {
	switch vv := v.(type) {
	case string:
		fmt.Println("is a string", vv)
	case float64:
		fmt.Println("is a float64", vv)
	case []interface{}:
		fmt.Println("is an array", vv)
		for k, v := range vv {
			fmt.Println(k, " ")
			walkJSON(v)
		}
	case map[string]interface{}:
		fmt.Println("is an object", vv)
		for k, v := range vv {
			fmt.Println(k, " ")
			walkJSON(v)
		}
	default:
		fmt.Println("Type not defined")
	}
}
