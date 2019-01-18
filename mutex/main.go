package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	mapvar := make(map[int]int)

	mutex := &sync.Mutex{}

	var readOp, writeOp int64
	for r := 0; r < 100; r++ {
		go func() {

			rand.Seed(time.Now().UnixNano())
			key := rand.Intn(10)
			mutex.Lock()
			val := mapvar[key]
			fmt.Println("Read ", r, ":", val)
			mutex.Unlock()
			atomic.AddInt64(&readOp, 1)
			time.Sleep(time.Microsecond)
		}()
	}

	for w := 0; w < 100; w++ {
		go func() {

			rand.Seed(time.Now().UnixNano())
			key := rand.Intn(10)
			val := rand.Intn(100)
			mutex.Lock()
			mapvar[key] = val
			fmt.Println("Write ", w, " key", key, " : ", val)
			mutex.Unlock()
			atomic.AddInt64(&writeOp, 1)

			time.Sleep(time.Microsecond)
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("read op succ:", atomic.LoadInt64(&readOp), "\n write op succ:", atomic.LoadInt64(&writeOp))
	fmt.Println(mapvar)
}
