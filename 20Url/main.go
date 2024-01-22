package main

import (
	"fmt"
	"net/url"
)
const myurl = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=123"
func main() {
	fmt.Println("URLs in Golang")
	result,_ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port()) 
	fmt.Println(result.RawQuery)

	qparams := result.Query() 
	fmt.Printf("The type of the parsed query params is %T\n", qparams)
	
	fmt.Println(qparams["coursename"])

	for _,val := range qparams {
		fmt.Println("Params are:", val)
	}

	partsOfUrl := &url.URL{	
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=abc", 
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}