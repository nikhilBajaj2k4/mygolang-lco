package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to get request wala video")
	PerformGetRequest("http://localhost:8000/get")
}

func PerformGetRequest(url string){
	var myurl string = url
	fmt.Println(myurl)

	resp, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(content)
	fmt.Println(string(content)) 

	defer resp.Body.Close()

	fmt.Println("Status code is:", resp.Status)
	fmt.Println("Content Length is:", resp.ContentLength)

	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	fmt.Println(byteCount)
	fmt.Println(responseString.String())

}