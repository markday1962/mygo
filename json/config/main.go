package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type configuration struct {
	Enabled bool
	Path    string
	Servers []Server
}

type Server struct {
	Name string
	ID   string
}

func main() {
	file, _ := os.Open("conf.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error", err)
	}
	if conf.Enabled == true {
		fmt.Println(conf.Path)
	}
	for _, v := range conf.Servers {
		fmt.Println(v.Name, ":", v.ID)
	}
}
