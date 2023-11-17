package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Model
type Course struct {
	CourseID    string `json:"courseid"`
	CourseName  string `json:"coursename"`
	CoursePrice int    `json:"price"`
	Author      *Author
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//fake DB

var course []Course

// middleware or helper
func (c *Course) IsEmpty() bool {
	return c.CourseID == "" && c.CourseName == ""
}

func main() {
	fmt.Println("This is example of Api")
}

//controller

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello this is API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get All Courses"))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(course)
}
