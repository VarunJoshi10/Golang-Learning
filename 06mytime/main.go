package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("This is example of time module")
	fmt.Println("Current time is ", time.Now())
	//Formatting the time
	fmt.Println("Formatted Time:", time.Now().Format("02-01-2006  15:04:05 Monday"))

	//Create New Date
	createdDate := time.Date(2002, time.July, 8, 3, 30, 30, 0, time.Local)
	fmt.Println("New Date is ", createdDate)

	//Formatted new Date
	fmt.Println("Formatted Date: ", createdDate.Format("02-01-2006"))
}
