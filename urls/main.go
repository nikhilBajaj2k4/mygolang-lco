package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs"

func main() {
	fmt.Println("URL")

	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}

	fmt.Println("Scheme: ", result.Scheme)
	fmt.Println("Path: ", result.Path)
	fmt.Println("Host: ", result.Host)
	fmt.Println("Port: ", result.Port())
	fmt.Println("Raw Query: ", result.RawQuery)
}