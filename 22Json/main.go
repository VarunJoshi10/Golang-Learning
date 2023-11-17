package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` //this 3rd parameter is alias
	Price    int
	Platform string   `json:"website"`
	Status   bool     `json:"-"`              //here using "-" it hides this
	Tags     []string `json:"tags,omitempty"` //using omitempty does not shows null values
}

func main() {
	fmt.Println("This is example of Json")
	JsonEncoder()

}

func JsonEncoder() {
	myCourse := []course{
		{"Golang", 100, "Youtube.com", true, []string{"go"}},
		{"Python", 200, "Youtube.com", false, []string{"py", "Scripting"}},
		{"Java", 300, "Youtube.com", true, nil},
	}
	result, _ := json.MarshalIndent(myCourse, "", "\t")
	fmt.Println(string(result))
}
