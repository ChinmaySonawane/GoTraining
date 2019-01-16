package main

import "fmt"

func main()  {
	
	var i interface{}
	
	i = 42
	show(i)
	
	i="play"
	show(i)
}

func show(i interface{})  {
	fmt.Printf("\n%v\t%T",i,i)
}