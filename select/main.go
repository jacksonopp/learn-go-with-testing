package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(urls ...string) (winner string, err error) {
	return ConfigurableRacer(tenSecondTimeout, urls...)
}

func ConfigurableRacer(timeout time.Duration, urls ...string) (winner string, err error) {
	select {
	case <-ping(urls[0]):
		return urls[0], nil
	case <-ping(urls[1]):
		return urls[1], nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for urls to respond")
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func main() {
	fmt.Println("Hello")
}
