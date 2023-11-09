package main

import "fmt"

func main() {
	fmt.Println("This is example of Methods")
	stud1 := Student{"Varun", "varun@go.dev", 78, true}
	stud1.GetStatus()
	stud2 := Student{"Pratik", "pratik@go.dev", 61, false}
	stud2.GetStatus()

}

type Student struct {
	Name   string
	Email  string
	Marks  int
	Status bool
}

func (s Student) GetStatus() {
	fmt.Printf("Status of %v is %v\n", s.Name, s.Status)
}
