package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := http.DefaultClient

	req, err := http.NewRequest("GET", "https://httpbin.org/get", nil)
	if err != nil {
		log.Fatalln("Unable to create request:", err)
	}
	//pass the request to the client
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("Unable to pass response")
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
