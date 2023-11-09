package main

import "fmt"

func main() {
	fmt.Println("This is example of Functions")
	fmt.Println("Total is ", adder(1, 2, 6, 1))
}

func adder(value ...int) int {
	total := 0

	for _, val := range value {
		total += val
	}
	return total
}
