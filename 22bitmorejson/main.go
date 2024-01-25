package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"` // it does not show the password in the json
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Hello World")
	// EncodeJson()
	DecodeJSON()
}

func EncodeJson() {

	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abcd", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "abcd", []string{"web-dev", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "abcd", []string{"web-dev", "js"}},
		{"LCO DevOps", 299, "LearnCodeOnline.in", "abcd", []string{"web-dev", "js"}},
		{"LCO Kubernetes", 299, "LearnCodeOnline.in", "abcd", []string{"web-dev", "js"}},
	}

	// package this data as json
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJSON() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": [
				"web-dev", 
				"js"
		]
	}
	`)

	var lcocourse course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcocourse)
		fmt.Printf("%#v\n", lcocourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is %T\n", k, v, v)
	}
}
