package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("This is example of slices")

	//In slices we dont need to define size as arrays

	var days = []string{"Monday", "Tuesday", "Wednesday"}

	fmt.Printf("type of days is %T\n", days)
	fmt.Println("Length :", len(days))

	days = append(days, "Thusday", "Friday", "Saturday")
	fmt.Println(days)
	fmt.Println("Length :", len(days))

	days = append(days[1:4])
	fmt.Println("slicing", days)

	// create slice using make
	num := make([]int, 4)
	num[0] = 123
	num[1] = 345
	num[2] = 23
	num[3] = 27
	// num[4] = 66 This causes error because size is defined

	num = append(num, 678, 875, 111) // this can be used to add elements when size is defined
	fmt.Println(sort.IntsAreSorted(num))
	fmt.Println("Numbers :", num)
	sort.Ints(num)
	fmt.Println("Numbers after sorting:", num)

	fmt.Println(sort.IntsAreSorted(num))

}
