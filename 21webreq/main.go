package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const url string = "http://localhost:8000/get"

func main() {
	fmt.Println("This is example of Request like GET,POST")
	// PerformGetRequest(url)
	PerformPostJsonRequest()

}

func PerformGetRequest(url string) {
	response, err := http.Get(url)
	checkNilErr(err)
	defer response.Body.Close()
	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Content Length:", response.ContentLength)
	content, _ := ioutil.ReadAll(response.Body)
	//One way
	// fmt.Println(string(content))
	//Another way
	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	fmt.Println(byteCount)
	fmt.Println(responseString.String())
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	responseBody := strings.NewReader(`
		{
			"name":"Varun"
		}
	`)

	response, _ := http.Post(myurl, "application/json", responseBody)

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
