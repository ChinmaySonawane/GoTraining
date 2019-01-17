package main

import (
	"fmt"
)

func test() { //v chan bool) {
	fmt.Println("start")
	//v <- true
	defer fmt.Println("end")
}
func main() {
	//v := make(chan bool, 1)
	go test() //v)
	//<-v
}
