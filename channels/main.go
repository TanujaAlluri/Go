package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://golang.org",
	}

	for _, url := range urls {
		checkLink(url)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
		return
	} else {
		fmt.Println(link, "is up")
	}
}
