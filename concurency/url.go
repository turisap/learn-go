package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
)

func checkAndSaveBody(url string, wg *sync.WaitGroup) {
	fmt.Println("\n" + strings.Repeat("@", 20))
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("%s is Down", url)
	} else {
		// clean up
		// if not done, connection is not reused ->
		// memory leaks
		defer resp.Body.Close()

		fmt.Printf("Status Code for %s is %d\n", url, resp.StatusCode)

		if resp.StatusCode == 200 {
			bodyBytes, err := io.ReadAll(resp.Body)

			if err != nil {
				fmt.Println("Could not parse body resp for\n", url)
			}

			fileName := strings.Split(url, "//")[1] + ".txt"

			fmt.Printf("Writing response body to %s\n", fileName)

			err = os.WriteFile(fileName, bodyBytes, 0644)

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	fmt.Println("\n" + strings.Repeat("@", 20))
	wg.Done()
}

func main() {
	urls := []string{"https://www.golang.org", "https://www.ololosh.org", "https://www.google.com", "https://www.medium.com"}

	var wg sync.WaitGroup

	wg.Add(len(urls))

	for _, v := range urls {
		go checkAndSaveBody(v, &wg)
	}

	fmt.Println("number of goroutines", runtime.NumGoroutine())
	wg.Wait()
}
