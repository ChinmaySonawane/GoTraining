package main

import "fmt"

type I interface
{
	fill()
}

type bottle struct
{
	liquid string
	capacity int 
}

func (b *bottle) fill()  {

	fmt.Println("filling ", b.liquid , "capacity", b.capacity, "liter's")
}

func main()  {
	
	//var i *I
	a := bottle{"water",2}
	a.fill()
}