package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, url := range urls {
		go checkLink(url, c)
	}

	// for { // infinite loop
	// 	go checkLink(<-c, c)
	// }

	for l := range c { // infinite loop
		go func() {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
			/*
				warning message: loop variable l captured by func literal.
				This essentially means that we are referencing loop variable l
				in both main routine and inside child routine.

				--------------------------------
				address	|	value
				---------------------------------
				100			|	google.com(value of l) <--------------main routine
				101			|				...											|
				102			|				...											|---------child routine

				For example, first we invoked all 5 website's checklink func.
				We now got response, first from amazon, and we pushed the url to channel.
				Then the value of 'l' is now 'Amazon'.
				The main routine waits for 5 sec, and creates a new Amazon child routine.
				And the Amazon child routine is waiting on a blocking http request.

				Meanwhile, Google routine finished, and the url is pushed.
				The new value of 'l' is 'Google'.
				At this point, if we reference 'l' in Amazon routine, that was spun
				the value would be 'Google' instead of 'Amazon'.

				Since both main routine and child routine are pointing to the same address.
			*/
		}()
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
		c <- link
		return
	}
	fmt.Println(link, "is up")
	c <- link
}
