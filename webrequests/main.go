package main

import (
	"fmt"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Web Request")

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	
	fmt.Printf("%T\n", resp)
	resp.Body.Close() // reminder: Always close the request
}