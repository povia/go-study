package main

import (
	"fmt"
	"net/http"
	"time"
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

	// c 에 response 가 존재하는 동안에는 계속해서 for loop 가 작동한다.
	for l := range c {
		// func literal 을 통해 5초를 쉰 후 CheckLink 를 호출한다.
		go func(link string) {
			time.Sleep(5 * time.Second)
			CheckLink(link, c)
		}(l) // function literal 에 전달하기 위한 인자 (link 에 매핑됨)
	}
}

func CheckLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return

	}

	fmt.Println(link, "might be up!")
	c <- link
}
