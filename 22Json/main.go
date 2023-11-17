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
	// JsonEncoder()
	JsonDecoder()

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

func JsonDecoder() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "Golang",
		"Price": 100,
		"website": "Youtube.com",
		"tags": [
				"go"
		]
	}
	`)
	//check validity
	checkValid := json.Valid(jsonDataFromWeb)

	//One way
	var lcoCourse course
	if checkValid {
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON IS NOT VALID")
	}

	// Another Way
	var myData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myData)
	fmt.Println(myData)

	//Iterating Over Map
	for k, v := range myData {
		fmt.Printf("%v : %v Type:%T\n", k, v, v)
	}

}
