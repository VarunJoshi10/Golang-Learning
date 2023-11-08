package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "time"
)

func main() {
	fmt.Println("This is maths example")

	//Random number using math/rand
	// rand.Seed(time.Now().Unix())
	// fmt.Println("random number using math package is ", rand.Intn(5)+1)

	//Random number using crypto/rand
	num, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println("random number using crypto package is ", num)
}
