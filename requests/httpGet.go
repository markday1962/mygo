package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// get request with query arguments which are key=value pairs
	resp, err := http.Get("https://httpbin.org/get?search=something")
	if err != nil {
		log.Fatalln("Unable to get response:", err)
	} else {
		fmt.Printf("Response to request \n%v\n\n", resp)
		fmt.Printf("request type: %T\n\n", resp)
	}
	defer resp.Body.Close()
	//read into a byte slice of content
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Unable to read body:", err)
	}
	//convert the content byte slice into a string
	fmt.Println("Printing content of response \n")
	fmt.Println(string(content))

}
