package main

import (
	"fmt"
	"sync"
	"time"
)

func print(s string, v *sync.WaitGroup) {

	defer v.Done()
	for index := 0; index < 5; index++ {

		time.Sleep(1000 * time.Microsecond)
		fmt.Println(s, index)
	}
}

func main() {
	var v sync.WaitGroup
	v.Add(2)
	go print("so", &v)
	go print("on", &v)
	//print("to", v)
	fmt.Println("to")
	v.Wait()
}
