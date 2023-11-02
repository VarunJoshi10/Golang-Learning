package main

import "fmt"

const Pi float32 = 3.14 // this is constant and is public

func main() {
	// Strings
	var username string = "Varun"
	fmt.Println(username)
	fmt.Printf("Type is %T\n", username)

	//Boolean
	var isEligible bool = true
	fmt.Println(isEligible)
	fmt.Printf("Type is %T \n", isEligible)

	//Integer
	var smallint uint8 = 255
	fmt.Println(smallint)
	fmt.Printf("Type is %T \n", smallint)

	//Float
	var smallFloat float32 = 255.12345678
	fmt.Println(smallFloat)
	fmt.Printf("Type is %T \n", smallFloat)

	//Another way of declaring
	var numberOfUsers = 3000
	fmt.Println(numberOfUsers)
	fmt.Printf("Type is %T \n", numberOfUsers)

	// One more way of declaring
	accesskey := 123432

	fmt.Println("Access key is ", accesskey)

	fmt.Println("Value of pie is ", Pi)

}
