package main

import (
	"fmt"
	"net/http"
	"time"
)

type Result struct {
	Url      string
	Response int
	Err      error
}

func getUrl(url string, ch chan<- Result) {
	client := http.Client{Timeout: 2 * time.Second}
	fmt.Println(url)
	result := Result{
		Url: url,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Request error:", err)
		result.Err = err
		ch <- result
		return
	}

	httpResp, err := client.Do(req)

	if err != nil {
		fmt.Println("Request failure", err)
		result.Err = err
		ch <- result
		return
	}
	defer httpResp.Body.Close()
	result.Response = httpResp.StatusCode
	ch <- result

}

func main() {
	var s []Result
	url := "http://badurl:8080"
	ch := make(chan Result)

	for i := 1; i <= 10; i++ {
		go getUrl(url, ch)
	}

	for i := 1; i <= 10; i++ {
		z := <-ch
		s = append(s, z)
	}

	fmt.Println(s)

}
