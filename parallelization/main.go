package main

import (
	"fmt"
	"runtime"
)

func PrintNo(x int, c chan bool, num int) {
	//x := int(^int(0))
	for i := 0; i < 1000000; i++ {

		fmt.Println(x, num)
		x++
	}
	c <- true
}

func main() {

	var numCPU = runtime.NumCPU()
	//var i int64 = 0
	fmt.Println(numCPU, ^uint(0))
	c := make(chan bool, numCPU)
	for i := 0; i < numCPU; i++ {
		go PrintNo(i, c, i)
	}
	for i := 0; i < numCPU-1; i++ {
		<-c
	}
}
