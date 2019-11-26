package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	//for {
	//	go checkLink(<- c, c)
	//}
	// Equivalent for
	for l := range c {
		go func(ln string) {
			time.Sleep(4 * time.Second)
			checkLink(ln, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down!")
		c <- link
	} else {
		fmt.Println(link, " is up!")
		c <- link
	}
}
