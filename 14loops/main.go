package main

import "fmt"

func main() {
	fmt.Println("This is example of Loops")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thusday", "Friday"}

	days = append(days, "Saturday")

	//Normal Loops
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("Index is %v and day is %v\n", index, day)
	// }

	value := 1
	for value < 10 {
		if value == 5 {
			value++
			continue
		}

		if value == 7 {
			goto report
		}

		if value == 8 {
			break
		}
		fmt.Println(value)
		value++
	}

report:
	fmt.Println("This is label for goto example")
}
