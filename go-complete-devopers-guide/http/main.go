package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func (logWriter) Write(bs []byte) (int, error) {
	length := len(bs)
	fmt.Println(string(bs))

	fmt.Println("Just wrote this many bytes:", length)
	return length, nil
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}
