package main

import "fmt"

func main() {
	nikhil := Person{"Nikhil", 16, true, 25, "nbajaj133@gmail.com"}

	fmt.Println(nikhil.Status)
	nikhil.GetTotal()
	nikhil.GetEmail()
}

type Person struct {
	Name   string
	Age    int
	Status bool
	Score  int
	Email  string
}

func (u Person) GetTotal() {
	fmt.Println(u.Age + u.Score)
}

func (u Person) GetEmail() {
	fmt.Println(u.Email)
}
