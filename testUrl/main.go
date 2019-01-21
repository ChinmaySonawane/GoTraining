package main

import (
	"Training/urlStatus"
	//	"fmt"
)

func main() {
	var done = make(chan bool)
	var urls = []string{"http://www.google.com/", "http://www.chinmay.com/", "http://www.chinmaysonawane.com/", "https://golang.org/doc/effective_go.html"}
	urlStatus.CheckUrls(urls, 2, done)
	<-done
	//fmt.println(x)
}
