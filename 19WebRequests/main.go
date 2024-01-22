package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Web Requests in Golang")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	// fmt.Println("Printing Response", response)
	// fmt.Println("Status Code:", response.StatusCode)
	// fmt.Println("Status:", response.Status)
	// fmt.Println("Header:", response.Header)
	databytes,err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println("Content:", content)
}
