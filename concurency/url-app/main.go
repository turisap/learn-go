package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func checkUrl(url string, c chan string) {
	resp, err := http.Get(url)

	if err != nil {

		s := fmt.Sprintf("%s is DOWN!\n", url)
		s += fmt.Sprintf("Error: %v  \n", err)

		fmt.Println(s)

		c <- url
	} else {
		defer resp.Body.Close()

		s := fmt.Sprintf("Status Code: %d  \n", resp.StatusCode)
		s += fmt.Sprintf("%s is UP\n", url)

		fmt.Println(s)

		c <- url
	}
}

func main() {
	urls := []string{"https://www.golang.org", "https://www.ololosh.org", "https://www.google.com", "https://www.medium.com"}
	ch := make(chan string)

	for _, url := range urls {
		go checkUrl(url, ch)
	}

	fmt.Println("Go routines n:", runtime.NumGoroutine())

	//for {
	//	go checkUrl(<-ch, ch)
	//	fmt.Println("\n" + strings.Repeat("@", 20))
	//	time.Sleep(time.Second * 2)
	//}

	//for url := range ch {
	//	time.Sleep(time.Second * 2)
	//	go checkUrl(url, ch)
	//}

	for url := range ch {
		go func(u string, c chan string) {
			time.Sleep(time.Second * 2)
			checkUrl(u, c)
		}(url, ch)
	}
}
