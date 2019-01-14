package main

import "fmt"

func main() {
	
	if i ,v:= 101,111; i == 10 && v == 11 {

		fmt.Println("in if")
	}else if v == 11 || i == 10 {
		
		fmt.Println("in else if")
	} else {
		
		fmt.Println("in else")
	}

}