package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("This is example of ifelse")

	fmt.Println("Enter any number")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	num, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	if num%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	fmt.Println("Enter another number ")
	input2, _ := reader.ReadString('\n')
	num2, _ := strconv.ParseInt(strings.TrimSpace(input2), 10, 64)

	fmt.Printf("Sum of %v and %v is %v\n", num, num2, num+num2)
}
