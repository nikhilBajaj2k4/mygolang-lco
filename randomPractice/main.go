package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome")
	content := "Ye ek piece of content he, writing to be specific"
	
	file, err := os.Create("./file.txt")
	errorCheck(err)

	length, err := io.WriteString(file, content)
	errorCheck(err)

	fmt.Println("Length is ", length)
	defer file.Close()
	readFile("./file.txt")
} 
 
func readFile(filename string){
	databyte, err := ioutil.ReadFile(filename)
	errorCheck(err)
	fmt.Println("Content of the file is \n", string(databyte))
}
 
func errorCheck(err error){
	if err != nil {
		panic(err)
	}
}
