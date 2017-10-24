// httpGet project main.go
package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	var status []int
	counter := make(map[int]int)

	var threads int = 100
	url := "http://localhost:8000/"
	for i := 0; i < threads; i++ {

		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			res, _ := http.Get(url)
			status = append(status, res.StatusCode)
			counter[res.StatusCode]++
		}(url)

	}
	wg.Wait()

	fmt.Printf("Threads %d\n", threads)
	fmt.Println("Result\tCount")
	for i := range counter {
		fmt.Println(i, "\t", counter[i])
	}

}
