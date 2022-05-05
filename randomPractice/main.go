package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome")
	content := "Ye ek piece of content he, writing to be specific"
	
	file, err := os.Create("./file.txt")
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("Length is ", length)
	defer file.Close()
}
 