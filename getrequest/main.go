package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to get request wala video")
	// PerformGetRequest("http://localhost:8000/get")
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest(url string){
	var myurl string = url
	fmt.Println(myurl)

	resp, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(content)
	fmt.Println(string(content)) 

	defer resp.Body.Close()

	fmt.Println("Status code is:", resp.Status)
	fmt.Println("Content Length is:", resp.ContentLength)

	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	fmt.Println(byteCount)
	fmt.Println(responseString.String())
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	//fake json payload

	requestBody := strings.NewReader(`
		{
			"coursename":"Let's go go goa with golang",
			"price": 0,
			"platform":"learnCodeOnline.in"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//formdata

	data := url.Values{}
	data.Add("firstname", "hitesh")
	data.Add("lastname", "choudhary")
	data.Add("email", "hitesh@goa.dev")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}