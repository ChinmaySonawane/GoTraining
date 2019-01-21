package main

import (
	"Training/test/website"
	"fmt"
)

func main() {
	var urls = []string{"http://www.google.com/", "http://www.chinmay.com/", "http://www.chinmaysonawane.com/", "https://golang.org/", "https://www.joshsoftware.com"}
	var done = make(chan *website.Website) //channel to pass CheckStatus func
	for _, url := range urls {
		webobj := website.NewWebsite(url)

		go webobj.CheckStatus(done) //making of goroutines
	}
	for range urls {
		webobj := <-done //channel receiver
		if webobj.Status == true {
			fmt.Println(webobj.Url, "is working")
			continue
		}
		fmt.Println(webobj.Url, "is down")
	}
}
