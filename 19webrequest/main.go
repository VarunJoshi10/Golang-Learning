package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("This is example of Web Request")
	resp, err := http.Get(url)
	checkNilErr(err)
	fmt.Printf("Type of Response is %T\n", resp)
	defer resp.Body.Close() //Always close the connection
	databyte, err := ioutil.ReadAll(resp.Body)
	checkNilErr(err)
	content := string(databyte)
	fmt.Println(content)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
