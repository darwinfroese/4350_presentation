package main

import (
	"net/http"
	"net/http/cookiejar"
)

func main() {
	cookieJar, _ := cookiejar.New(nil)

	httpClient := &http.Client{
		CheckRedirect: nil,
		Jar:           cookieJar,
	}

	// Send your HttpRequests
	httpClient, _ := http.NewRequest("GET", url, nil)
}
