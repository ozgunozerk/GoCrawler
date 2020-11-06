package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func crawler(wg *sync.WaitGroup, urlChannel <-chan string) {

	defer wg.Done()
	client := &http.Client{Timeout: 15 * time.Second} // single client is sufficient for multiple requests

	for urlItem := range urlChannel {

		req1, _ := http.NewRequest("GET", "http://"+urlItem, nil)                                           // generating the request
		req1.Header.Add("User-agent", "Mozilla/5.0 (X11; Linux i586; rv:31.0) Gecko/20100101 Firefox/74.0") // changing user-agent
		resp1, respErr1 := client.Do(req1)                                                                  // sending the prepared request and getting the response
		if respErr1 != nil {
			fmt.Println("server error", urlItem)
			continue
		}

		if resp1.StatusCode/100 == 2 { // means server responded with 2xx code

			f1, fileErr1 := os.Create("200/" + urlItem + "_original.txt") // creating the relative file
			if fileErr1 != nil {
				fmt.Println("file error", urlItem)
				log.Fatal(fileErr1)
			}

			_, writeErr1 := io.Copy(f1, resp1.Body) // writing the sourcecode into our file
			if writeErr1 != nil {
				fmt.Println("file error", urlItem)
				continue
			}
			f1.Close()
			resp1.Body.Close()

			fmt.Println("success:", urlItem)

		}
	}
}

func main() {

	var wg sync.WaitGroup // synchronization to wait for all the goroutines

	file, err := os.Open("urls.txt") // the file containing the url's
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // don't forget to close the file

	urlChannel := make(chan string) // create a channel to store all the url's

	_ = os.Mkdir("200", 0755) // if it's there, it will create an error, and we will simply ignore it

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go crawler(&wg, urlChannel)
	}

	scanner := bufio.NewScanner(file) // each line has another url
	for scanner.Scan() {
		urlChannel <- scanner.Text()
	}
	close(urlChannel)
	wg.Wait()
}
