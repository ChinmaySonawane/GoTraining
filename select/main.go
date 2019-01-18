package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int, 1)
	c2 := make(chan int, 1)

	go func() {
		time.Sleep(2 * time.Microsecond)
		c1 <- 0
	}()

	go func() {
		time.Sleep(time.Microsecond)
		c2 <- 1
	}()

	for i := 0; i < 3; i++ {

		select {
		case m := <-c1:
			fmt.Println(m)
		case n := <-c2:
			fmt.Println(n)
		default:
			fmt.Println("default")
		}
		time.Sleep(2 * time.Microsecond)
	}
}
