package main

import "fmt"

func main() {
	fmt.Println("This is example of pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	mynum := 8
	var ptr = &mynum //& is used for referencing
	fmt.Println("Actual value of pointer is ", ptr)
	fmt.Println("Actual value of pointer is ", *ptr) //* is used for actual value

	*ptr = *ptr + 2
	fmt.Println("Actual value of pointer is ", ptr)
	fmt.Println("Actual value of pointer is ", *ptr)
}
