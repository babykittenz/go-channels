package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}

func checkMultipleLinks(links []string) {
	c := make(chan string) // we make a channel to communicate between goroutines
	// create a goroutine for each link
	for _, link := range links {
		go checkLink(link, c) // we pass the channel to the goroutine
	}

	// we wait for the goroutines to finish
	for l := range c {
		go func(link string) { // we use a parameter to avoid calling a variable that is outside the scope of the goroutine
			time.Sleep(5 * time.Second) // we wait 5 seconds before checking the link again so that we don't spam the server
			checkLink(link, c)
		}(l) // we call the goroutine with l as link
	}
}
