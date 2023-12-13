package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start go routine
	}

	var results []string
	for range os.Args[1:] {
		result := <-ch
		fmt.Println(result)
		results = append(results, result)
	}
	writeFile(results)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to the channel
		return
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func writeFile(values []string) {
	fmt.Println("Creating users.json file...")
	f, err := os.Create(fmt.Sprintf("users.json"))
	if err != nil {
		fmt.Printf("Error while creating the file: %s\n", err.Error())
	}

	fmt.Println("Writing file...")
	for _, value := range values {
		_, err = f.WriteString(fmt.Sprintf("%s\n", value))
		if err != nil {
			fmt.Printf("Error while writing to the file: %s\n", err.Error())
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Printf("Error while defer f: %s\n", err.Error())
	}
}
