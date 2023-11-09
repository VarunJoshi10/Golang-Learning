package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("This is example of Switch Case")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("Number is 1 now you can open")
	case 2:
		fmt.Println("Number is 2")
	case 3:
		fmt.Println("Number is 3")
		fallthrough //This can be used to run next case
	case 4:
		fmt.Println("Number is 4")
	case 5:
		fmt.Println("Number is 5")
	case 6:
		fmt.Println("Number is 6")
	default:
		fmt.Println("Something went wrong")
	}
}
