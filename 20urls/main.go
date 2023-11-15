package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/login?coursename=reactjs&userid=vj123"

func main() {
	fmt.Println("This is example of Urls")
	fmt.Println(myurl)
	//Parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	for _, val := range qparams {
		fmt.Println("Param is ", val)
	}

	partOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "ValorantFlashCard",
		Path:     "Agents",
		RawQuery: "Agentid=10&AgentName=Raze",
	}

	newUrl := partOfUrl.String()
	fmt.Println(newUrl)
	res, _ := url.Parse(newUrl)
	fmt.Println(res.RawQuery)
	qpara := res.Query()
	for _, val := range qpara {
		fmt.Println("Param:", val)
	}
}
