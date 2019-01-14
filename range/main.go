package main

import "fmt"

func main()  {
	var str = "abcd"
	for i,j := range str {
		fmt.Printf("\n %v %T %v %T", i,i, string(j),j)
	}
}