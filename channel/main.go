package main

import (
	"fmt"
)

func check(x int, c chan int) {
	fmt.Println("check", x)
	c <- x
	c <- x + 90
	//c <- x + 2
	//fmt.Println(<-c)
	//fmt.Println(<-c)
}

func main() {
	c := make(chan int, 1)
	fmt.Println("main", c)
	//c <- 1
	//c <- 2
	//w := sync.WaitGroup
	//w.add(1)
	go func() {

		for i := 0; i < 10; i++ {
			check(i, c)
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println("1", <-c)
		fmt.Println("2", <-c)
	}
}
