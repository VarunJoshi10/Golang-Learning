package main

import "fmt"

func main() {
	fmt.Println("This is a example of Structs")
	// No inheritance in Golang
	varun := User{"Varun", "varun@go.dev", true, 21}

	fmt.Println("User details are: ", varun)
	fmt.Printf("%+v\n", varun)
	fmt.Printf("Name is %v Email is %v\n", varun.Name, varun.Email)
}

// First letter is capital to make it public
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
