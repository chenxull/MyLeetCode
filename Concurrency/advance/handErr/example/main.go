package main

import (
	"fmt"
	"net/http"
)

type Result struct {
	Error    error
	Response *http.Response
}

func main() {
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result)

		go func() {
			defer close(results)
			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{Error: err, Response: resp} //3
				select {
				case <-done:
					return
				case results <- result:
				}
			}
		}()
		return results
	}

	done := make(chan interface{})
	defer close(done)
	errCount := 0
	urls := []string{"https://www.baidu.com", "https://badhost", "s", "d", "r", "f"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil { //5
			fmt.Printf("error: %v\n", result.Error)
			errCount++
			if errCount >= 3 {
				fmt.Println("too many errors,breaking!")
				break
			}
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}
