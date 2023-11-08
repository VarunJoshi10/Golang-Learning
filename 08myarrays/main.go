package main

import "fmt"

func main() {
	fmt.Println("This is example of Arrays")

	//Declaration
	var names [4]string

	//Initilization
	names[0] = "Varun"
	names[1] = "Pratik"
	names[2] = "Aryan"
	names[3] = "Manish"
	fmt.Println("Name list contains :", names)
	fmt.Println("Length of name list :", len(names))

	// when there is no element it adds space
	// Single file declaration and initialization
	var lang = [4]string{"Golang", "Python", "Bash"}

	fmt.Println("Languages :", lang)
	fmt.Println("Length of languages list :", len(lang))
}
