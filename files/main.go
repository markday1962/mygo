package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const (
	S3B      = "aistemos-cloud-init"
	S3PREFIX = "/swarm-config/"

	FDIR = "/etc/cipher-conf/"
)

func main() {
	files := listDir(FDIR)
	copyObjToS3(files)
}

// Copy Objects to a S3 bucket
func copyObjToS3(files []string) {
	tb := S3B + S3PREFIX + getHostname() + "/"
	fmt.Println("The target bucket and prefix key is", tb)
	for _, f := range files {
		fmt.Println(tb + f)
	}
}

// Gets the os hostname
func getHostname() string {
	hname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	return hname
}

// Returns a *slice of string containing names of files in a directory
func listDir(d string) []string {
	filesx := []string{}
	files, err := ioutil.ReadDir(d)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		fmt.Println(f.Name())
		filesx = append(filesx, f.Name())
	}
	return filesx
}
