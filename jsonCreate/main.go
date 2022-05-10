package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Age      int
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags     []string `json:"tags,omitempty"` //if the value is null, dont show it
}

func main() {
	lcoCourses := []course{
		{"React Course", 150, "Udemy", "menahibatauga", []string{"webdev", "javascript"}},
		{"Angular Course", 170, "Udemy", "mekabhibatauga", []string{"google", "b"}},
		{"Ret Course", 120, "Udemy", "meshayadbatauga", nil},
	}
	finalJson, _ := json.MarshalIndent(lcoCourses, "", "\t")
	fmt.Printf("%s\n",finalJson)
}