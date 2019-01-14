package main

import "fmt"

func a() {
	
	defer b(c())
	fmt.Println("A")
}

func b(x int) {

	fmt.Println("B ",x)
}

func c()int   {
	fmt.Println("C")
	return 10
}

func main()  {
	
	defer fmt.Println("Main")
	a()
}