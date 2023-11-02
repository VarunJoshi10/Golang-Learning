package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to User Input")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Your Lucky Number")

	// comma ok, comma err syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Luck number is :", input)
	fmt.Printf("Type of input is %T", input)
}
