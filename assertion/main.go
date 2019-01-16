package main

import "fmt"

func main()  {
	
	var i interface{}="ha"

	//v := i.(float64)
	//fmt.Println(v)

	v, ok := i.(float64)
	fmt.Println(v, ok)

	//v, ok := i.(int64)
	fmt.Println(v, ok)

	f,ok1 := i.(string)
	fmt.Println(f ,ok1)
}