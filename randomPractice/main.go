package main

import "fmt"

func main() {
	defer fmt.Println("Pehla")
	defer fmt.Println("Doosra")
	defer fmt.Println("Teesra")

	fmt.Println("Wan")
	MyDefer()
	// Wan, 43210, Teesra, Doosra, Pehla
}

func MyDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
