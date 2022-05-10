package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Web Request")

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("Response Type is: %T\n", resp)
	defer resp.Body.Close() // reminder: Always close the request

	databytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)
}