package main

import ("fmt"
		"reflect"
)
type node struct{
	x, y int
}

func (p node)add()int  {
	fmt.Println(reflect.TypeOf(p))
	return p.x+p.y
}

func main()  {
	
	a := node{3,4}
	fmt.Println(a.add())
	fmt.Println((&a).add())
}