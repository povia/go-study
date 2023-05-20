package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)
	for _, link := range links {
		go CheckLink(link, c)
	}

	fmt.Println(<-c)
}

func CheckLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think"
		os.Exit(1)

	}
	c <- "Yep its up"
	fmt.Println(link, "might be up!")
}
