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
	for _, link := range links {
		go CheckLink(link)
	}
}

func CheckLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		os.Exit(1)
	}
	fmt.Println(link, "might be up!")
}
