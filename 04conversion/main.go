package main

import "fmt"

func main() {
	nikhil := Person{"Nikhil", 16, true, 25}

	fmt.Println(nikhil.Status)
	nikhil.GetTotal()
}

type Person struct {
	Name   string
	Age    int
	Status bool
	Score  int
}

func (u Person) GetTotal() {
	fmt.Println(u.Age + u.Score)
}
