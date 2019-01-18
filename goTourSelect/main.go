package main

import "fmt"

func fibonacci(c, quit chan int) {
	x := 0
	for {
		select {
		case c <- x:
			fmt.Println("case c")
			x++
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go fibonacci(c, quit)
	func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
}
