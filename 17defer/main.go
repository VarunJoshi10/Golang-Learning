package main

import "fmt"

func main() {
	fmt.Println("This is example of Defer")
	//using defer puts that line at then end for execution
	//Line is pushed into the stack
	defer fmt.Println("Statement 1") // stack=["statement1"]
	defer fmt.Println("Statement 2") // stack=["statement1","statement2"]
	defer fmt.Println("Statement 3") // stack=["statement1","statement2","statement3"]

	fmt.Println("Line before method")  //This line will execute First
	example()                          //function call execution
	fmt.Println("After before method") // This line will execute normally
	//Now elements from main stack will pop
	// statement3, statement2, statement1
}

func example() {
	//in function elements will be in another stack
	for i := 0; i < 5; i++ {
		defer fmt.Println(i) // stack=[0,1,2,3,4]
	}
	//Elements will pop from stack
	// 4,3,2,1,0 then back to main
}
