package main

import (
	"fmt"
	// "time"
	"net/http"
	"sync"
)

var wg sync.WaitGroup 
func main() {
	fmt.Println("Hello, playground")
	// go greeter("Hello")
	// greeter("Gopher")

	webSiteList := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.amazon.com",
		"https://www.golang.org",
		"https://www.stackoverflow.com",
		"https://www.reddit.com",
		"https://www.wikipedia.org",
		"https://www.theguardian.com",
	}

	for _, url := range webSiteList {
		go getStatusCode(url)
		wg.Add(1)
	}
	wg.Wait()
}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(5 * time.Second)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(url string) {
	defer wg.Done()
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("OOPS error")
	}
	defer res.Body.Close()
	fmt.Printf("%s is up and running with status code: %d\n", url, res.StatusCode)
}
