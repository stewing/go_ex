package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSeconds = 10 * time.Second

func Racer(url1, url2 string) (string, error) {
	return ConfigurableRacer(url1, url2, tenSeconds)
}

func ConfigurableRacer(url1, url2 string, timeout time.Duration) (string, error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", url1, url2)
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
