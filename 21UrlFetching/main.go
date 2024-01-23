package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	// PerformGetRequest()
	PerformPostRequest()
}

func PerformGetRequest() {
	const url = "http://localhost:8080"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Status Code : ", resp.StatusCode)
	fmt.Println("Response length : ", resp.ContentLength)

	var responseString strings.Builder

	// to read the response body
	content, _ := ioutil.ReadAll(resp.Body)
	bytecontent, _ := responseString.Write(content)
	fmt.Println("Response Body:", bytecontent)

	// to write the response body to a file

	// fmt.Println("Response Body : ", responseString.String())
	// fmt.Println("Response Body : ", string(content))
}

func PerformPostRequest() {
	const url = "http://localhost:8000/post"
	content1 := strings.NewReader(`{"name":"John"}`)
	resp, err := http.Post(url, "application/json", content1)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	} 
	fmt.Println("Response Body:", string(content))
}
